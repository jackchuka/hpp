package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientGet_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("key") == "" {
			t.Fatal("expected key param")
		}
		if r.URL.Query().Get("format") != "json" {
			t.Fatal("expected format=json")
		}
		_, _ = w.Write([]byte(`{"results":{"results_available":1}}`))
	}))
	defer srv.Close()

	c := NewClient("test-key")
	c.BaseURL = srv.URL

	var result struct {
		Results struct {
			ResultsAvailable int `json:"results_available"`
		} `json:"results"`
	}
	err := c.Get("/gourmet/v1/", nil, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Results.ResultsAvailable != 1 {
		t.Fatalf("expected 1, got %d", result.Results.ResultsAvailable)
	}
}

func TestClientGet_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"results":{"error":[{"message":"Invalid API key","code":2000}]}}`))
	}))
	defer srv.Close()

	c := NewClient("bad-key")
	c.BaseURL = srv.URL

	var result struct{}
	err := c.Get("/gourmet/v1/", nil, &result)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Code != 2000 {
		t.Fatalf("expected code 2000, got %d", apiErr.Code)
	}
}
