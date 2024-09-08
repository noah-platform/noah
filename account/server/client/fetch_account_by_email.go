package client

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	"github.com/noah-platform/noah/account/server/handler"
	"github.com/noah-platform/noah/pkg/httputil"
	"github.com/noah-platform/noah/pkg/response"
)

type GetAccountByEmailResponse = handler.GetAccountByEmailResponse

func (c *Client) FetchAccountByEmail(email string) (*GetAccountByEmailResponse, error) {
	path := fmt.Sprintf("/internal/v1/accounts/email/%s", url.PathEscape(email))

	account, status, err := httputil.Get[response.DataResponse[GetAccountByEmailResponse]](c.client, path)
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
