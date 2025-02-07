package client

import (
	"github.com/pkg/errors"

	"github.com/noah-platform/noah/auth/server-session/handler"
	"github.com/noah-platform/noah/pkg/httputil"
	"github.com/noah-platform/noah/pkg/response"
)

type CreateSessionRequest = handler.CreateSessionRequest

type CreateSessionResponse = handler.CreateSessionResponse

func (c *Client) CreateSession(userID, ipAddress, userAgent string) (*CreateSessionResponse, error) {
	path := "/internal/v1/sessions"

	session, _, err := httputil.Post[response.DataResponse[CreateSessionResponse]](c.client, path, &CreateSessionRequest{
		UserID:    userID,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create session")
	}

	return &session.Data, nil
}
