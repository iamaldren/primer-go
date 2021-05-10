package primer

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
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

func (c *Client) get(ctx context.Context, idempotencyKey, path string, apiReq interface{}, apiResp DTO) error {
	req, err := c.createRequest(http.MethodGet, idempotencyKey, path, apiReq)
	if err != nil {
		return wrapError(err)
	}
	return c.do(ctx, req, apiResp)
}

func (c *Client) post(ctx context.Context, idempotencyKey, path string, apiReq interface{}, apiResp DTO) error {
	req, err := c.createRequest(http.MethodPost, idempotencyKey, path, apiReq)
	if err != nil {
		return wrapError(err)
	}

	return c.do(ctx, req, apiResp)
}

func (c *Client) createRequest(method, idempotencyKey, path string, request interface{}) (*http.Request, error) {
	body, err := marshalRequest(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("X-Idempotency-Key", idempotencyKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, apiResp DTO) error {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return wrapError(err)
	}
	defer resp.Body.Close()
	return decodeResponse(resp, apiResp)
}

func marshalRequest(request interface{}) (io.Reader, error) {
	if request == nil {
		return nil, nil
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func decodeResponse(resp *http.Response, apiResp DTO) error {
	requestID := resp.Header.Get("X-Grabkit-Grab-Requestid")
	apiResp.SetRequestID(requestID)

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(apiResp); err != nil {
			return wrapError(err)
		}
		return nil
	case http.StatusNoContent:
		return nil
	default:
		var msg string
		if resp.ContentLength != 0 {
			bb, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return wrapError(err)
			}
			msg = string(bb)
		}
		return &Error{
			Status:    resp.StatusCode,
			Message:   msg,
			RequestID: requestID,
		}
	}
}
