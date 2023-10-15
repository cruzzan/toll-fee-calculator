package dayMatcher

import "time"

// WeekdayMatcher matches instants that fall on the configured day of the week. For example a saturday.
type WeekdayMatcher struct {
	day time.Weekday
}

func (m WeekdayMatcher) Matches(date time.Time) bool {
	return date.Weekday() == m.day
}

func NewWeekdayMatcher(day time.Weekday) WeekdayMatcher {
	return WeekdayMatcher{day: day}
}
