package velora

import (
	"encoding/json"
	"time"
)

const (
	StreamOnline  = "stream.online"
	StreamOffline = "stream.offline"
	StreamUpdate  = "stream.update"
)

type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Timestamp time.Time       `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}
