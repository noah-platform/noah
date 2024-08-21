package client

import "github.com/noah-platform/noah/pkg/httputil"

type Config = httputil.Config

type Client struct {
	client *httputil.Client
}

func New(cfg Config) *Client {
	return &Client{
		client: httputil.New(cfg),
	}
}
