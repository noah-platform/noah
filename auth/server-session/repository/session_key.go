package repository

import (
	"fmt"
	"strings"
)

func redisSessionKey(id string) string {
	return fmt.Sprintf("session:%s", strings.TrimSpace(id))
}
