package dayMatcher

import "time"

// DateRangeMatcher matches an instant that falls between (Inclusive) the two configured time.Time instants
type DateRangeMatcher struct {
	start time.Time
	end   time.Time
}

func (m DateRangeMatcher) Matches(date time.Time) bool {
	return (date.Equal(m.start) || date.After(m.start)) && (date.Equal(m.end) || date.Before(m.end))
}

func NewDateRangeMatcher(start, end time.Time) DateRangeMatcher {
	return DateRangeMatcher{start, end}
}
