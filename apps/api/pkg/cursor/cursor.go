package cursor

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

// RecordCursor encodes the position of the last seen record for stable keyset pagination.
// Primary sort key: recorded_at DESC; tiebreaker: id DESC.
type RecordCursor struct {
	RecordedAt time.Time `json:"recorded_at"`
	ID         string    `json:"id"`
}

// Encode serialises a RecordCursor to a URL-safe base64 string.
func Encode(c RecordCursor) (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Decode parses a cursor string produced by Encode.
// Returns nil when the input is empty (first page).
func Decode(s string) (*RecordCursor, error) {
	if s == "" {
		return nil, nil
	}
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	var c RecordCursor
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
