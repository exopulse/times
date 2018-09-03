package times

import "time"

// Parse parses time using format auto-detection. At the moment RFC3339 format is assumed.
func Parse(value string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", value)
}
