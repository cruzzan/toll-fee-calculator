package dayMatcher

import (
	"testing"
	"time"
)

func TestWeekdayMatcher_Matches(t *testing.T) {
	type testCase struct {
		in   time.Time
		want bool
	}

	testCases := []testCase{
		{time.Date(2023, 10, 11, 23, 59, 59, 0, time.UTC), false},
		{time.Date(2023, 10, 12, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2023, 10, 13, 13, 37, 0, 0, time.UTC), false},
		{time.Date(1999, 1, 1, 13, 37, 0, 0, time.UTC), false},
		{time.Date(1998, 12, 31, 13, 37, 0, 0, time.UTC), true},
	}

	m := NewWeekdayMatcher(time.Thursday)
	for i, c := range testCases {
		if m.Matches(c.in) != c.want {
			t.Fatalf("Matcher failed expectation for test case %d, %s wanted %t", i, c.in, c.want)
		}
	}
}
