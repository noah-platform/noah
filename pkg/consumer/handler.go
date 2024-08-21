package consumer

import (
	"context"
	"encoding/json"

	"github.com/noah-platform/noah/pkg/messaging"
)

type Handler interface {
	Handle(ctx context.Context, event messaging.Event, payload json.RawMessage) error
}
