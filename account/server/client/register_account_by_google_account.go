package client

import (
	"github.com/pkg/errors"

	"github.com/noah-platform/noah/account/server/handler"
	"github.com/noah-platform/noah/pkg/httputil"
	"github.com/noah-platform/noah/pkg/response"
)

type RegisterAccountByGoogleAccountRequest = handler.RegisterAccountByGoogleAccountRequest

type RegisterAccountByGoogleAccountResponse = handler.RegisterAccountByGoogleAccountResponse

func (c *Client) RegisterAccountByGoogleAccount(email, name, googleAccountID string) (*RegisterAccountByGoogleAccountResponse, error) {
	path := "/internal/v1/accounts/google"

	account, status, err := httputil.Post[response.DataResponse[RegisterAccountByGoogleAccountResponse]](c.client, path, &RegisterAccountByGoogleAccountRequest{
		Email:           email,
		Name:            name,
		GoogleAccountID: googleAccountID,
	})
	if err != nil {
		switch status {
		case 409:
			return nil, ErrAccountAlreadyExists
		default:
			return nil, errors.Wrap(err, "failed to register account")
		}
	}

	return &account.Data, nil
}
