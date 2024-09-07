package client

import (
	"github.com/pkg/errors"

	"github.com/noah-platform/noah/account/server/handler"
	"github.com/noah-platform/noah/pkg/httputil"
	"github.com/noah-platform/noah/pkg/response"
)

type GetAccountByEmailResponse = handler.GetAccountByEmailResponse

func (c *Client) FetchAccountByEmail(email string) (*GetAccountByEmailResponse, error) {
	path := "/internal/v1/accounts"

	type query struct {
		Email string `url:"email"`
	}
	account, status, err := httputil.GetWithQuery[response.DataResponse[GetAccountByEmailResponse]](c.client, path, query{
		Email: email,
	})
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
