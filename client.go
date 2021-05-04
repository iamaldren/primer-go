package primer

import (
	"net/http"
	"strings"
)

// Client will be used to make requests to Primer APIs
type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// DTO ...
type DTO interface {
	SetRequestID(id string)
}

// ClientOption constructor parameter for NewClient(...)
type ClientOption func(*Client) error

// NewClient constructs a new Client which can make requests to the Primer APIs.
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	if strings.TrimSpace(c.apiKey) == "" {
		return nil, ErrCredentialsMissing
	}
	if strings.TrimSpace(c.baseURL) == "" {
		return nil, ErrBaseURLMissing
	}
	return c, nil
}

// WithHTTPClient configures a Primer API client with a http.Client to make requests over.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		if c.Transport == nil {
			c.Transport = http.DefaultTransport
		}
		client.httpClient = c
		return nil
	}
}

// WithAPIKey configures a Primer API client with an API Key
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) error {
		c.apiKey = apiKey
		return nil
	}
}

// WithBaseURL configures a Primer API client with a custom base url
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}
