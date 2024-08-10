package messaging

import (
	"encoding/json"
	"time"
)

type TraceID string

const TraceIDContextKey TraceID = "traceId"

type Message struct {
	TraceID   string          `json:"traceId" validate:"required"`
	Event     Event           `json:"event" validate:"required"`
	Payload   json.RawMessage `json:"payload" validate:"required"`
	Timestamp time.Time       `json:"timestamp" validate:"required"`
}
