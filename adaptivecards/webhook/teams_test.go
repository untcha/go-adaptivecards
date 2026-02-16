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
	if err := PostToWorkflowRaw(context.Background(), srv.URL, card); err != nil {
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
	err := PostToWorkflowRaw(context.Background(), srv.URL, card)
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
	if !strings.Contains(err.Error(), "context canceled") {
		t.Fatalf("expected context canceled error, got %v", err)
	}
}

func TestPostToWorkflowRawWithClientSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	card := c.NewCard().AddTextBlock(els.NewTextBlock("hello"))
	if err := PostToWorkflowRawWithClient(context.Background(), srv.Client(), srv.URL, card); err != nil {
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
	err := PostToWorkflowRaw(context.Background(), srv.URL, invalid)
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
	err := PostToWorkflowRaw(context.Background(), srv.URL, card)
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

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
