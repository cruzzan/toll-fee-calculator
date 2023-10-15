package dayMatcher

import "time"

// DateMatcher matches if the instant occurs on the same day as one of the configured dates.
// This is for example used to hold holiday dates
type DateMatcher struct {
	dates []time.Time
}

func (m DateMatcher) Matches(date time.Time) bool {
	needle := date.Truncate(24 * time.Hour)

	for _, h := range m.dates {
		if h.Equal(needle) {
			return true
		}
	}

	return false
}

func NewDateMatcher(dates []time.Time) DateMatcher {
	return DateMatcher{dates: dates}
}
