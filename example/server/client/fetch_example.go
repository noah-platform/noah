package client

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	"github.com/noah-platform/noah/example/server/handler"
	"github.com/noah-platform/noah/pkg/httputil"
	"github.com/noah-platform/noah/pkg/response"
)

type GetExampleResponse = handler.GetExampleResponse

func (c *Client) FetchExample(exampleID string) (*GetExampleResponse, error) {
	path := fmt.Sprintf("/internal/v1/example/%s", url.PathEscape(exampleID))

	example, status, err := httputil.Get[response.DataResponse[GetExampleResponse]](c.client, path)
	if err != nil {
		switch status {
		case 404:
			return nil, ErrExampleNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch example")
		}
	}

	return &example.Data, nil
}
