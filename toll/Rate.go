package toll

import (
	"fmt"
	"time"
)

type Rate struct {
	start time.Time
	end   time.Time
	price int
}

// Matches checks if the rate applies to the given instant
func (r Rate) Matches(instant time.Time) bool {
	start := time.Date(instant.Year(), instant.Month(), instant.Day(), r.start.Hour(), r.start.Minute(), r.start.Second(), 0, instant.Location())
	end := time.Date(instant.Year(), instant.Month(), instant.Day(), r.end.Hour(), r.end.Minute(), r.end.Second(), 0, instant.Location())

	if r.start.Equal(r.end) {
		return instant.Equal(start)
	}

	return (instant.Equal(start) || instant.After(start)) && (instant.Equal(end) || instant.Before(end))
}

func (r Rate) GetPrice() int {
	return r.price
}

func NewRate(start, end string, price int) (Rate, error) {
	s, err := parseTimeFromString(start, false)
	if err != nil {
		return Rate{}, fmt.Errorf("error occured while creating rate with start '%s': %w", start, err)
	}

	e, err := parseTimeFromString(end, true)
	if err != nil {
		return Rate{}, fmt.Errorf("error occured while creating rate with end '%s': %w", end, err)
	}

	return Rate{s, e, price}, nil
}

func parseTimeFromString(ts string, max bool) (time.Time, error) {
	format := "15:04"
	t, err := time.Parse(format, ts)
	if err != nil {
		return time.Time{}, fmt.Errorf("error occured while parsing '%s' to time: %w", ts, err)
	}

	// For rate end we want to max the minute out.
	if max {
		t = t.Add(1 * time.Minute)
		t = t.Add(-1 * time.Nanosecond)
	}

	return t, nil
}
