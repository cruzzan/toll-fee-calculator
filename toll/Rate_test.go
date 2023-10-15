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

	for i, c := range cases {
		format := "15:04"
		start, err := time.Parse(format, c.s)
		if err != nil {
			t.Errorf("Error occurred while parsing test case data index %d, start %s could not be parsed to time", i, c.s)
		}
		end, err := time.Parse(format, c.e)
		if err != nil {
			t.Errorf("Error occurred while parsing test case data index %d, end %s could not be parsed to time", i, c.e)
		}
		r := Rate{start, end, 0}

		if r.Matches(c.i) != c.want {
			t.Fatalf("Expected %t result when mtaching %s against %s -> %s", c.want, c.i, r.start, r.end)
		}
	}
}
