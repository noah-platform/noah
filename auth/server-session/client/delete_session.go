package client

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	"github.com/noah-platform/noah/pkg/httputil"
)

func (c *Client) DeleteSession(sessionID string) error {
	path := fmt.Sprintf("/internal/v1/sessions/%s", url.PathEscape(sessionID))

	_, status, err := httputil.Delete[struct{}](c.client, path)
	if err != nil {
		switch status {
		case 404:
			return ErrSessionNotFound
		default:
			return errors.Wrap(err, "failed to register account")
		}
	}

	return nil
}
