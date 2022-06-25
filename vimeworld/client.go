package vimeworld

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
)

const (
	defaultBaseURL = "https://api.vimeworld.com/"
)

// Client for VimeWorld API.
type Client struct {
	client  *http.Client
	baseURL *url.URL

	lastToken uint64
	tokens    []string
}

// Options for Client.
type Options struct {
	Client  *http.Client
	BaseURL string
	Tokens  []string
}

func (o *Options) setDefaults() {
	if o.Client == nil {
		o.Client = http.DefaultClient
	}
	if o.BaseURL == "" {
		o.BaseURL = defaultBaseURL
	}
}

// NewClient returns new vimeworld client.
// Parameters in the options are optional.
func NewClient(opts Options) (*Client, error) {
	opts.setDefaults()

	parsedURL, err := url.Parse(opts.BaseURL)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(parsedURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", parsedURL)
	}

	return &Client{
		client:  opts.Client,
		baseURL: parsedURL,
		tokens:  opts.Tokens,
	}, nil
}

func (c *Client) getToken() string {
	n := atomic.AddUint64(&c.lastToken, 1)
	return c.tokens[(int(n)-1)%len(c.tokens)]
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	if len(c.tokens) > 0 {
		q := u.Query()
		q.Add("token", c.getToken())

		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// BareDo sends an API request and lets you handle the api response.
func (c *Client) BareDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("ctx must be not nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	return resp, err
}

// ErrorResponse struct.
type ErrorResponse struct {
	Error Error `json:"error"`
}

// Do sends an API request and returns the API response.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	var errResponse ErrorResponse

	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(rawData, &errResponse)

	if errResponse.Error.ErrorCode != 0 {
		return nil, &errResponse.Error
	}

	if err := json.Unmarshal(rawData, &v); err != nil {
		return nil, err
	}

	return resp, nil
}
