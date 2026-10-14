package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

const defaultBaseURL = "https://webservice.recruit.co.jp/hotpepper"

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("hotpepper API error %d: %s", e.Code, e.Message)
}

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: defaultBaseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type errorResponse struct {
	Results struct {
		Error []APIError `json:"error"`
	} `json:"results"`
}

func (c *Client) Get(path string, params interface{}, out interface{}) error {
	u := c.BaseURL + path

	vals, err := query.Values(params)
	if err != nil {
		return fmt.Errorf("encoding params: %w", err)
	}
	vals.Set("format", "json")

	// The API only accepts the key as a query parameter; keep it out of error messages.
	req, err := http.NewRequest("GET", u+"?"+vals.Encode(), nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	vals.Set("key", c.APIKey)
	req.URL.RawQuery = vals.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		var urlErr *url.Error
		if errors.As(err, &urlErr) {
			urlErr.URL = redactKey(urlErr.URL)
		}
		return fmt.Errorf("making request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	// Check for API-level errors first
	var raw json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	var errResp errorResponse
	if err := json.Unmarshal(raw, &errResp); err == nil && len(errResp.Results.Error) > 0 {
		e := errResp.Results.Error[0]
		return &APIError{Code: e.Code, Message: e.Message}
	}

	return json.Unmarshal(raw, out)
}

func redactKey(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "[invalid URL]"
	}
	q := u.Query()
	if q.Has("key") {
		q.Set("key", "REDACTED")
		u.RawQuery = q.Encode()
	}
	return u.String()
}
