package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	c "github.com/untcha/go-adaptivecards/adaptivecards/card"
)

const (
	defaultHTTPTimeout = 15 * time.Second
	maxErrorBodyBytes  = 8 << 10 // 8 KiB
)

// URLPolicy controls security checks for workflow URLs.
// Defaults are intentionally strict to reduce SSRF risk.
type URLPolicy struct {
	// AllowHTTP allows plain HTTP URLs. Default: false (HTTPS only).
	AllowHTTP bool
	// AllowPrivateNetworks allows loopback/private/link-local targets. Default: false.
	AllowPrivateNetworks bool
	// AllowedHosts optionally restricts outbound requests to exact hostnames.
	// If empty, no hostname allowlist is applied.
	AllowedHosts []string
}

func defaultURLPolicy() URLPolicy {
	return URLPolicy{
		AllowHTTP:            false,
		AllowPrivateNetworks: false,
		AllowedHosts:         nil,
	}
}

// PostToWorkflowRaw posts the Adaptive Card as the WHOLE request body (no wrapper).
func PostToWorkflowRaw(ctx context.Context, url string, card *c.Card) error {
	return PostToWorkflowRawWithClientAndPolicy(
		ctx,
		&http.Client{Timeout: defaultHTTPTimeout},
		url,
		card,
		defaultURLPolicy(),
	)
}

// PostToWorkflowRawWithClient posts the Adaptive Card as the whole request
// body, using the provided HTTP client.
func PostToWorkflowRawWithClient(ctx context.Context, client *http.Client, url string, card *c.Card) error {
	return PostToWorkflowRawWithClientAndPolicy(ctx, client, url, card, defaultURLPolicy())
}

// PostToWorkflowRawWithClientAndPolicy posts the Adaptive Card as the whole request
// body, using the provided HTTP client and URL security policy.
func PostToWorkflowRawWithClientAndPolicy(
	ctx context.Context,
	client *http.Client,
	url string,
	card *c.Card,
	policy URLPolicy,
) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if client == nil {
		return errors.New("http client is nil")
	}
	if url == "" {
		return errors.New("workflow URL is empty")
	}
	if card == nil {
		return errors.New("card is nil")
	}
	if err := validateWorkflowURL(ctx, url, policy); err != nil {
		return err
	}

	// Validate the card first
	if err := card.Validate(); err != nil { // ensure it's valid before POST
		return err
	}

	b, err := json.Marshal(card) // Card.MarshalJSON injects defaults
	if err != nil {
		return fmt.Errorf("marshal card: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// #nosec G704 -- URL is validated by validateWorkflowURL before issuing the request.
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("post workflow: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		slurp, readErr := io.ReadAll(io.LimitReader(resp.Body, maxErrorBodyBytes))
		if readErr != nil {
			return fmt.Errorf("workflow HTTP %d: read error response: %w", resp.StatusCode, readErr)
		}
		if len(slurp) == 0 {
			return fmt.Errorf("workflow HTTP %d", resp.StatusCode)
		}
		return fmt.Errorf("workflow HTTP %d: %s", resp.StatusCode, string(slurp))
	}

	return nil
}

func validateWorkflowURL(ctx context.Context, rawURL string, policy URLPolicy) error {
	u, err := parseWorkflowURL(rawURL)
	if err != nil {
		return err
	}

	if err := validateWorkflowScheme(u, policy); err != nil {
		return err
	}

	hostname := strings.ToLower(u.Hostname())
	if err := validateWorkflowAllowedHost(hostname, policy); err != nil {
		return err
	}

	if policy.AllowPrivateNetworks {
		return nil
	}

	// Prevent direct private/loopback/link-local IP targets.
	if ip := net.ParseIP(hostname); ip != nil {
		if isPrivateLikeIP(ip) {
			return fmt.Errorf("workflow URL host %q resolves to a private or local address", hostname)
		}
		return nil
	}

	// Prevent DNS-based SSRF to private/local networks.
	ips, err := net.DefaultResolver.LookupIPAddr(ctx, hostname)
	if err != nil {
		return fmt.Errorf("resolve workflow URL host %q: %w", hostname, err)
	}
	for _, ipAddr := range ips {
		if isPrivateLikeIP(ipAddr.IP) {
			return fmt.Errorf(
				"workflow URL host %q resolves to private or local address %q",
				hostname,
				ipAddr.IP.String(),
			)
		}
	}
	return nil
}

func parseWorkflowURL(rawURL string) (*url.URL, error) {
	s := strings.TrimSpace(rawURL)
	if s == "" {
		return nil, errors.New("workflow URL is empty")
	}
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("invalid workflow URL: %w", err)
	}
	if !u.IsAbs() {
		return nil, errors.New("workflow URL must be absolute")
	}
	if u.User != nil {
		return nil, errors.New("workflow URL must not contain userinfo")
	}
	if u.Hostname() == "" {
		return nil, errors.New("workflow URL host is empty")
	}
	return u, nil
}

func validateWorkflowScheme(u *url.URL, policy URLPolicy) error {
	scheme := strings.ToLower(u.Scheme)
	if scheme == "https" {
		return nil
	}
	if scheme == "http" && policy.AllowHTTP {
		return nil
	}
	if policy.AllowHTTP {
		return fmt.Errorf("workflow URL scheme must be http or https (got %q)", u.Scheme)
	}
	return fmt.Errorf("workflow URL scheme must be https (got %q)", u.Scheme)
}

func validateWorkflowAllowedHost(hostname string, policy URLPolicy) error {
	if len(policy.AllowedHosts) == 0 {
		return nil
	}
	allowed := make([]string, 0, len(policy.AllowedHosts))
	for _, h := range policy.AllowedHosts {
		n := strings.ToLower(strings.TrimSpace(h))
		if n != "" {
			allowed = append(allowed, n)
		}
	}
	if len(allowed) == 0 {
		return nil
	}
	if !slices.Contains(allowed, hostname) {
		return fmt.Errorf("workflow URL host %q is not in allowed hosts list", hostname)
	}
	return nil
}

func isPrivateLikeIP(ip net.IP) bool {
	return ip.IsLoopback() ||
		ip.IsPrivate() ||
		ip.IsLinkLocalUnicast() ||
		ip.IsLinkLocalMulticast() ||
		ip.IsMulticast() ||
		ip.IsInterfaceLocalMulticast() ||
		ip.IsUnspecified()
}
