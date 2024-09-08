package httputil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	qs "github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Config struct {
	BaseURL  string
	RetryMax int
}

type Client struct {
	config Config
	client *http.Client
}

func New(cfg Config) *Client {
	client := retryablehttp.NewClient()
	client.RetryMax = cfg.RetryMax

	return &Client{
		config: cfg,
		client: client.StandardClient(),
	}
}

func Get[R any](c *Client, path string) (*R, int, error) {
	return request[R](c, path, struct{}{})
}

func GetWithQuery[R, Q any](c *Client, path string, query Q) (*R, int, error) {
	return request[R](c, path, query)
}

func request[R, Q any](c *Client, path string, query Q) (*R, int, error) {
	q, err := qs.Values(query)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to parse query values")
	}

	url, err := url.JoinPath(c.config.BaseURL, path)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to join path")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to create request")
	}

	req.URL.RawQuery = q.Encode()

	resp, status, err := do[R](c.client, req)
	if err != nil {
		return nil, status, errors.Wrap(err, "failed to do request")
	}

	return resp, status, nil
}

func Post[R, B any](c *Client, path string, body *B) (*R, int, error) {
	return requestWithBody[R](c, http.MethodPost, path, body, struct{}{})
}

func PostWithQuery[R, B, Q any](c *Client, path string, body *B, query Q) (*R, int, error) {
	return requestWithBody[R](c, http.MethodPost, path, body, query)
}

func requestWithBody[R, B, Q any](c *Client, method string, path string, body *B, query Q) (*R, int, error) {
	q, err := qs.Values(query)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to parse query values")
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to marshal body")
	}

	url, err := url.JoinPath(c.config.BaseURL, path)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to join path")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to create request")
	}

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	resp, status, err := do[R](c.client, req)
	if err != nil {
		return nil, status, errors.Wrap(err, "failed to do request")
	}

	return resp, status, nil

}

func do[R any](c *http.Client, req *http.Request) (*R, int, error) {
	r, err := c.Do(req)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to do request")
	}
	defer r.Body.Close()

	if r.StatusCode < http.StatusOK || r.StatusCode >= http.StatusBadRequest {
		log.Debug().
			Str("method", req.Method).
			Str("url", req.URL.String()).
			Int("status", r.StatusCode).
			Str("requestId", r.Header.Get("X-Request-Id")).
			Msg("[httputil.do] client received unexpected status code")

		return nil, r.StatusCode, errors.Errorf("unexpected status code: %d", r.StatusCode)
	}

	var resp R
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, r.StatusCode, errors.Wrap(err, "failed to decode response")
	}

	return &resp, r.StatusCode, nil
}
