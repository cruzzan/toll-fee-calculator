package toll

import (
	"time"
)

type Rate struct {
	start time.Time
	end   time.Time
	price int
}

// Matches checks if the rate applies to the given instant
func (r Rate) Matches(instant time.Time) bool {
	h, m, s := instant.Clock()
	t := time.Date(0, 1, 1, h, m, s, 0, time.UTC)

	if r.start.Equal(r.end) {
		return t.Equal(r.start)
	}

	return (t.Equal(r.start) || t.After(r.start)) && (t.Equal(r.end) || t.Before(r.end))
}

func (r Rate) GetPrice() int {
	return r.price
}

func NewRate(start, end string, price int) Rate {
	format := "15:04"
	s, _ := time.Parse(format, start)
	e, _ := time.Parse(format, end)
	return Rate{s, e, price}
}
