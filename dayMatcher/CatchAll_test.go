package dayMatcher

import (
	"testing"
	"time"
)

func TestCatchAll_Matches(t *testing.T) {
	testCases := []time.Time{
		time.Now(),
		time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2300, 7, 13, 10, 15, 0, 0, time.UTC),
	}

	m := CatchAll{}
	for i, c := range testCases {
		if !m.Matches(c) {
			t.Fatalf("Catch all should match any date, case %d did not get a match", i)
		}
	}
}
