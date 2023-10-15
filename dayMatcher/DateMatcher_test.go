package dayMatcher

import (
	"testing"
	"time"
)

func TestDateMatcher_Matches(t *testing.T) {
	type testCase struct {
		in   time.Time
		want bool
	}

	testCases := []testCase{
		{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2023, 1, 1, 13, 37, 0, 0, time.UTC), true},
		{time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC), false},
		{time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC), true},
		{time.Date(2023, 10, 14, 13, 42, 0, 0, time.UTC), true},
		{time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC), false},
		{time.Date(2023, 10, 13, 23, 59, 59, 0, time.UTC), false},
	}

	m := NewDateMatcher([]time.Time{
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC),
	})
	for i, c := range testCases {
		if m.Matches(c.in) != c.want {
			t.Fatalf("Matcher failed expectation for test case %d, wanted %t ", i, c.want)
		}
	}
}
