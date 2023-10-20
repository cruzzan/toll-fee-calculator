package main

import (
	"testing"
	"time"
	"toll-fee-calculator/dayMatcher"
	"toll-fee-calculator/toll"
)

func TestCalculator_GetTollFee(t *testing.T) {
	type testCase struct {
		name string
		v    Vehicle
		p    []time.Time
		want int
	}
	testCases := []testCase{
		{
			"Motorbike should be toll free",
			Motorbike{},
			[]time.Time{
				time.Date(2023, 10, 5, 6, 45, 0, 0, time.UTC),
			},
			0,
		}, {
			"Car single passage",
			Car{},
			[]time.Time{
				time.Date(2023, 10, 5, 6, 45, 0, 0, time.UTC),
			},
			16,
		}, {
			"Car max toll",
			Car{},
			[]time.Time{
				time.Date(2023, 10, 5, 6, 31, 0, 0, time.UTC), // 16
				time.Date(2023, 10, 5, 7, 35, 0, 0, time.UTC), // 22
				time.Date(2023, 10, 5, 8, 36, 0, 0, time.UTC), // 9
				time.Date(2023, 10, 5, 15, 5, 0, 0, time.UTC), // 16
				time.Date(2023, 10, 5, 16, 6, 0, 0, time.UTC), // 22
				time.Date(2023, 10, 5, 17, 8, 0, 0, time.UTC), // 16
			},
			60,
		}, {
			"Car multiple passages in grace period",
			Car{},
			[]time.Time{
				time.Date(2023, 10, 5, 6, 31, 0, 0, time.UTC),
				time.Date(2023, 10, 5, 6, 32, 0, 0, time.UTC),
				time.Date(2023, 10, 5, 6, 33, 0, 0, time.UTC),
				time.Date(2023, 10, 5, 6, 34, 0, 0, time.UTC),
			},
			16,
		}, {
			"Car morning commute",
			Car{},
			[]time.Time{
				time.Date(2023, 10, 5, 7, 3, 0, 0, time.UTC),
				time.Date(2023, 10, 5, 7, 13, 0, 0, time.UTC),
				time.Date(2023, 10, 5, 7, 33, 0, 0, time.UTC),
				time.Date(2023, 10, 5, 8, 2, 0, 0, time.UTC),
				time.Date(2023, 10, 5, 8, 5, 0, 0, time.UTC),
				time.Date(2023, 10, 5, 8, 15, 0, 0, time.UTC),
			},
			38,
		},
	}

	calc := NewCalculator(getTestConfig())

	for _, c := range testCases {
		got := calc.GetTollFee(c.v, c.p...)
		if got != c.want {
			t.Fatalf("Test on case '%s' failed, wanted %d but got %d", c.name, c.want, got)
		}
	}
}

func getTestConfig() []toll.Config {
	var config []toll.Config

	config = append(config, toll.NewConfig(
		dayMatcher.CatchAll{},
		setUpTestRate("00:00", "05:59", 0),
		setUpTestRate("06:00", "06:29", 9),
		setUpTestRate("06:30", "06:59", 16),
		setUpTestRate("07:00", "07:59", 22),
		setUpTestRate("08:00", "08:29", 16),
		setUpTestRate("08:30", "14:59", 9),
		setUpTestRate("15:00", "15:29", 16),
		setUpTestRate("15:30", "16:59", 22),
		setUpTestRate("17:00", "17:59", 16),
		setUpTestRate("18:00", "18:29", 9),
		setUpTestRate("18:30", "23:59", 0),
	))

	return config
}

func setUpTestRate(start, end string, price int) toll.Rate {
	rate, _ := toll.NewRate(start, end, price)

	return rate
}
