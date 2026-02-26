package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	vals.Set("key", c.APIKey)
	vals.Set("format", "json")

	req, err := http.NewRequest("GET", u+"?"+vals.Encode(), nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
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
