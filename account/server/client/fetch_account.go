package client

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	"github.com/noah-platform/noah/account/server/handler"
	"github.com/noah-platform/noah/pkg/httputil"
	"github.com/noah-platform/noah/pkg/response"
)

type GetAccountResponse = handler.GetAccountResponse

func (c *Client) FetchAccount(userID string) (*GetAccountResponse, error) {
	path := fmt.Sprintf("/internal/v1/accounts/%s", url.PathEscape(userID))

	account, status, err := httputil.Get[response.DataResponse[GetAccountResponse]](c.client, path)
	if err != nil {
		switch status {
		case 404:
			return nil, ErrAccountNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch account")
		}
	}

	return &account.Data, nil
}
