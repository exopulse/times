package times

import (
	"log"
	"time"
)

// Parse parses time using format auto-detection.
func Parse(value string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, value)

	if err != nil {
		return time.Parse("2006-01-02T15:04:05-0700", value)
	}

	return t, nil
}

// MustParse is a helper that wraps a call to Parse() function and panics if the error is non-nil.
func MustParse(value string) time.Time {
	time, e := Parse(value)

	if e != nil {
		log.Panicf("failed to parse time [%s]: %s", value, e)
	}

	return time
}
