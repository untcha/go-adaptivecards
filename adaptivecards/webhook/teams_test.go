package webhook

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	c "github.com/untcha/go-adaptivecards/adaptivecards/card"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

func TestPostToWorkflowRawNilCard(t *testing.T) {
	err := PostToWorkflowRaw(context.Background(), "https://example.com", nil)
	if err == nil {
		t.Fatalf("expected error for nil card")
	}
	if !strings.Contains(err.Error(), "card is nil") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPostToWorkflowRawSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST method")
		}
		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Fatalf("unexpected content type: %s", got)
		}
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()

	card := c.NewCard().AddTextBlock(els.NewTextBlock("hello"))
	policy := URLPolicy{AllowHTTP: true, AllowPrivateNetworks: true}
	if err := PostToWorkflowRawWithClientAndPolicy(context.Background(), srv.Client(), srv.URL, card, policy); err != nil {
		t.Fatalf("unexpected post error: %v", err)
	}
}

func TestPostToWorkflowRawStatusError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad request"))
	}))
	defer srv.Close()

	card := c.NewCard().AddTextBlock(els.NewTextBlock("hello"))
	policy := URLPolicy{AllowHTTP: true, AllowPrivateNetworks: true}
	err := PostToWorkflowRawWithClientAndPolicy(context.Background(), srv.Client(), srv.URL, card, policy)
	if err == nil {
		t.Fatalf("expected non-2xx error")
	}
	if !strings.Contains(err.Error(), "workflow HTTP 400") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPostToWorkflowRawEmptyURL(t *testing.T) {
	err := PostToWorkflowRaw(context.TODO(), "", c.NewCard().AddTextBlock(els.NewTextBlock("x")))
	if err == nil || !strings.Contains(err.Error(), "workflow URL is empty") {
		t.Fatalf("expected empty URL error, got %v", err)
	}
}

func TestPostToWorkflowRawWithCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := PostToWorkflowRaw(ctx, "https://example.com", c.NewCard().AddTextBlock(els.NewTextBlock("x")))
	if err == nil {
		t.Fatalf("expected canceled context error")
	}
	if !strings.Contains(err.Error(), "canceled") {
		t.Fatalf("expected canceled error, got %v", err)
	}
}

func TestPostToWorkflowRawWithClientSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	card := c.NewCard().AddTextBlock(els.NewTextBlock("hello"))
	policy := URLPolicy{AllowHTTP: true, AllowPrivateNetworks: true}
	if err := PostToWorkflowRawWithClientAndPolicy(context.Background(), srv.Client(), srv.URL, card, policy); err != nil {
		t.Fatalf("unexpected post error: %v", err)
	}
}

func TestPostToWorkflowRawWithClientNilClient(t *testing.T) {
	err := PostToWorkflowRawWithClient(
		context.Background(),
		nil,
		"https://example.com",
		c.NewCard().AddTextBlock(els.NewTextBlock("x")),
	)
	if err == nil || !strings.Contains(err.Error(), "http client is nil") {
		t.Fatalf("expected nil client error, got %v", err)
	}
}

func TestPostToWorkflowRawInvalidCard(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	invalid := c.NewCard().AddTextBlock(els.NewTextBlock(""))
	policy := URLPolicy{AllowHTTP: true, AllowPrivateNetworks: true}
	err := PostToWorkflowRawWithClientAndPolicy(context.Background(), srv.Client(), srv.URL, invalid, policy)
	if err == nil {
		t.Fatalf("expected validation error for invalid card")
	}
	if called {
		t.Fatalf("server should not be called for invalid card")
	}
}

func TestPostToWorkflowRawStatusErrorBodyIsLimited(t *testing.T) {
	veryLargeBody := strings.Repeat("x", maxErrorBodyBytes+4096)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(veryLargeBody))
	}))
	defer srv.Close()

	card := c.NewCard().AddTextBlock(els.NewTextBlock("hello"))
	policy := URLPolicy{AllowHTTP: true, AllowPrivateNetworks: true}
	err := PostToWorkflowRawWithClientAndPolicy(context.Background(), srv.Client(), srv.URL, card, policy)
	if err == nil {
		t.Fatalf("expected non-2xx error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "workflow HTTP 400") {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(msg) > len("workflow HTTP 400: ")+maxErrorBodyBytes {
		t.Fatalf("expected limited error body, got length %d", len(msg))
	}
}

func TestPostToWorkflowRawWithClientTransportError(t *testing.T) {
	badClient := &http.Client{
		Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) {
			return nil, errors.New("network down")
		}),
	}
	card := c.NewCard().AddTextBlock(els.NewTextBlock("hello"))
	err := PostToWorkflowRawWithClient(context.Background(), badClient, "https://example.com", card)
	if err == nil || !strings.Contains(err.Error(), "post workflow") {
		t.Fatalf("expected wrapped transport error, got %v", err)
	}
}

func TestPostToWorkflowRawRejectsInsecureOrLocalURLByDefault(t *testing.T) {
	card := c.NewCard().AddTextBlock(els.NewTextBlock("hello"))

	err := PostToWorkflowRaw(context.Background(), "http://example.com/webhook", card)
	if err == nil || !strings.Contains(err.Error(), "must be https") {
		t.Fatalf("expected https enforcement error, got %v", err)
	}

	err = PostToWorkflowRaw(context.Background(), "https://127.0.0.1/webhook", card)
	if err == nil || !strings.Contains(err.Error(), "private or local") {
		t.Fatalf("expected local address rejection, got %v", err)
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
