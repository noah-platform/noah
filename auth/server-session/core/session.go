package core

import (
	"net/netip"
	"time"
)

type Session struct {
	SessionID string     `json:"sessionId"`
	UserID    string     `json:"userId"`
	IPAddress netip.Addr `json:"ipAddress"`
	UserAgent string     `json:"userAgent"`
	CreatedAt time.Time  `json:"createdAt"`
}
