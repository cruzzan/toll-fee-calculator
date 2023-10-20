package toll

import (
	"testing"
	"time"
)

func TestRate_Matches(t *testing.T) {
	type testCase struct {
		s    string
		e    string
		i    time.Time
		want bool
	}
	cases := []testCase{
		{"16:00", "22:00", time.Date(2023, 10, 14, 19, 47, 0, 0, time.UTC), true},
		{"22:12", "22:12", time.Date(2023, 10, 14, 19, 47, 0, 0, time.UTC), false},
		{"21:21", "21:21", time.Date(2023, 8, 10, 21, 21, 0, 0, time.UTC), true},
		{"21:00", "21:21", time.Date(2023, 8, 10, 21, 0, 0, 0, time.UTC), true},
		{"21:00", "21:21", time.Date(2023, 8, 10, 20, 59, 59, 0, time.UTC), false},
		{"21:00", "21:21", time.Date(2023, 8, 10, 21, 21, 0, 0, time.UTC), true},
		{"21:00", "21:21", time.Date(2023, 8, 10, 21, 22, 0, 0, time.UTC), false},
	}

	for _, c := range cases {
		r, err := NewRate(c.s, c.e, 0)
		if err != nil {
			t.Errorf("Error occured while setting up Rate %s", err)
		}

		if r.Matches(c.i) != c.want {
			t.Fatalf("Expected %t result when matching %s against %s -> %s", c.want, c.i, c.s, c.e)
		}
	}
}

func TestParseTimeFromString(t *testing.T) {
	type testCase struct {
		in        string
		max       bool
		want      time.Time
		wantError bool
	}

	testCases := []testCase{
		{"01:01", false, time.Date(0, 1, 1, 1, 1, 0, 0, time.UTC), false},
		{"20:13", false, time.Date(0, 1, 1, 20, 13, 0, 0, time.UTC), false},
		{"01:01", true, time.Date(0, 1, 1, 1, 1, 59, 999999999, time.UTC), false},
		{"20:13", true, time.Date(0, 1, 1, 20, 13, 59, 999999999, time.UTC), false},
		{"25:01", false, time.Time{}, true},
		{"10:61", false, time.Time{}, true},
		{"10:10:10", false, time.Time{}, true},
		{"2023-10-20 13:37:00", false, time.Time{}, true},
		{"hello", false, time.Time{}, true},
	}

	for _, c := range testCases {
		gotTime, gotErr := parseTimeFromString(c.in, c.max)

		if !c.want.Equal(gotTime) {
			t.Errorf("Wanted %s but got %s", c.want, gotTime)
		}

		if (gotErr != nil) != c.wantError {
			t.Fatalf("Expected error state %t, but got %t", c.wantError, gotErr != nil)
		}
	}
}
