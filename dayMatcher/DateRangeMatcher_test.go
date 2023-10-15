package dayMatcher

import (
	"testing"
	"time"
)

func TestDateRangeMatcher_Matches(t *testing.T) {
	type testCase struct {
		in   time.Time
		want bool
	}

	testCases := []testCase{
		{time.Date(2023, 3, 31, 23, 59, 59, 0, time.UTC), false},
		{time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2023, 4, 5, 13, 37, 0, 0, time.UTC), true},
		{time.Date(2023, 4, 9, 23, 59, 59, 0, time.UTC), true},
		{time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2023, 4, 10, 23, 59, 49, 0, time.UTC), true},
		{time.Date(2023, 4, 11, 0, 0, 0, 0, time.UTC), false},
	}

	m := NewDateRangeMatcher(
		time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 10, 23, 59, 59, 0, time.UTC),
	)
	for i, c := range testCases {
		if m.Matches(c.in) != c.want {
			t.Fatalf("Matcher failed expectation for test case %d, %s wanted %t", i, c.in, c.want)
		}
	}
}
