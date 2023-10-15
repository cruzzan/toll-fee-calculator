package main

import (
	"fmt"
	"time"
	"toll-fee-calculator/dayMatcher"
	"toll-fee-calculator/toll"
)

func main() {
	var config []toll.Config

	config = append(config, toll.NewConfig(
		dayMatcher.CatchAll{},
		toll.NewRate("00:00", "05:59", 0),
		toll.NewRate("06:00", "06:29", 9),
		toll.NewRate("06:30", "06:59", 16),
		toll.NewRate("07:00", "07:59", 22),
		toll.NewRate("08:00", "08:29", 16),
		toll.NewRate("08:30", "14:59", 9),
		toll.NewRate("15:00", "15:29", 16),
		toll.NewRate("15:30", "16:59", 22),
		toll.NewRate("17:00", "17:59", 16),
		toll.NewRate("18:00", "18:29", 9),
		toll.NewRate("18:30", "23:59", 0),
	))

	config = append(config, toll.NewConfig(
		dayMatcher.NewWeekdayMatcher(time.Saturday),
	))

	config = append(config, toll.NewConfig(
		dayMatcher.NewWeekdayMatcher(time.Sunday),
	))

	config = append(config, toll.NewConfig(
		dayMatcher.NewDateMatcher(get2023Holidays()),
	))

	config = append(config, toll.NewConfig(
		dayMatcher.NewDateMatcher(get2023DayBeforeHolidays()),
	))

	config = append(config, toll.NewConfig(
		dayMatcher.NewDateRangeMatcher(
			time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 7, 31, 23, 59, 59, 0, time.UTC),
		),
	))

	tollCalculator := NewCalculator(config)

	// Start serving requests to the calculator
	fmt.Printf("Passage costed %d\n", tollCalculator.GetTollFee(Car{}, time.Date(2023, 10, 13, 14, 15, 0, 0, time.UTC)))
}
