package dayMatcher

import "time"

// CatchAll type matches any given instant
type CatchAll struct {
}

func (m CatchAll) Matches(date time.Time) bool {
	return true
}
