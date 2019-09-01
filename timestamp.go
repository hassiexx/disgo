package disgo

import "time"

// Timestamp is the type for Discord ISO8601 timestamps.
// Timestamps are never parsed automatically because timestamps
// could be null which could cause errors or invalid conversions
// while unmarshaling a JSON.
type Timestamp string

// Parse parses the ISO8601 timestamp into a time.Time object.
func (t Timestamp) Parse() (time.Time, error) {
	return time.Parse(time.RFC3339Nano, string(t))
}
