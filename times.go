package times

import "time"

// Parse parses time using format auto-detection.
func Parse(value string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, value)

	if err != nil {
		return time.Parse("2006-01-02T15:04:05-0700", value)
	}

	return t, nil
}
