// Package dayMatcher contains a set of matchers that can be used to determine if an instant of time.Time falls
// within the matcher's config.
//
// They are used to determine if a set of rate rules should be applied to a passage instant
package dayMatcher

import "time"

type DayMatcher interface {
	Matches(date time.Time) bool
}
