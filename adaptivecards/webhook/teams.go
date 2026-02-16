package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	c "github.com/untcha/go-adaptivecards/adaptivecards/card"
)

const (
	defaultHTTPTimeout = 15 * time.Second
	maxErrorBodyBytes  = 8 << 10 // 8 KiB
)

// PostToWorkflowRaw posts the Adaptive Card as the WHOLE request body (no wrapper).
func PostToWorkflowRaw(ctx context.Context, url string, card *c.Card) error {
	return PostToWorkflowRawWithClient(ctx, &http.Client{Timeout: defaultHTTPTimeout}, url, card)
}

// PostToWorkflowRawWithClient posts the Adaptive Card as the whole request
// body, using the provided HTTP client.
func PostToWorkflowRawWithClient(ctx context.Context, client *http.Client, url string, card *c.Card) error {
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
