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

	config = append(config, toll.NewConfig(
		dayMatcher.NewWeekdayMatcher(time.Saturday),
	))

	config = append(config, toll.NewConfig(
		dayMatcher.NewWeekdayMatcher(time.Sunday),
	))

	config = append(config, toll.NewConfig(
		dayMatcher.CatchAll{},
		setUpRate("00:00", "05:59", 0),
		setUpRate("06:00", "06:29", 9),
		setUpRate("06:30", "06:59", 16),
		setUpRate("07:00", "07:59", 22),
		setUpRate("08:00", "08:29", 16),
		setUpRate("08:30", "14:59", 9),
		setUpRate("15:00", "15:29", 16),
		setUpRate("15:30", "16:59", 22),
		setUpRate("17:00", "17:59", 16),
		setUpRate("18:00", "18:29", 9),
		setUpRate("18:30", "23:59", 0),
	))

	tollCalculator := NewCalculator(config)

	// Start serving requests to the calculator
	fmt.Printf("Passage normal day costed %d\n", tollCalculator.GetTollFee(
		Car{},
		time.Date(2023, 10, 13, 14, 15, 0, 0, time.UTC),
		time.Date(2023, 10, 13, 15, 15, 10, 0, time.UTC),
		time.Date(2023, 10, 13, 16, 15, 20, 0, time.UTC),
	))

	fmt.Printf("Passage in july costed %d\n", tollCalculator.GetTollFee(
		Car{},
		time.Date(2023, 07, 13, 14, 15, 0, 0, time.UTC),
		time.Date(2023, 07, 13, 15, 15, 10, 0, time.UTC),
		time.Date(2023, 07, 13, 16, 15, 20, 0, time.UTC),
	))

	fmt.Printf("Passage on saturday costed %d\n", tollCalculator.GetTollFee(
		Car{},
		time.Date(2023, 10, 7, 14, 15, 0, 0, time.UTC),
		time.Date(2023, 10, 7, 15, 15, 10, 0, time.UTC),
		time.Date(2023, 10, 7, 16, 15, 20, 0, time.UTC),
	))

	fmt.Printf("Passage on sunday costed %d\n", tollCalculator.GetTollFee(
		Car{},
		time.Date(2023, 10, 8, 14, 15, 0, 0, time.UTC),
		time.Date(2023, 10, 8, 15, 15, 10, 0, time.UTC),
		time.Date(2023, 10, 8, 16, 15, 20, 0, time.UTC),
	))

	fmt.Printf("Passage on day before christmas costed %d\n", tollCalculator.GetTollFee(
		Car{},
		time.Date(2023, 12, 23, 14, 15, 0, 0, time.UTC),
		time.Date(2023, 12, 23, 15, 15, 10, 0, time.UTC),
		time.Date(2023, 12, 23, 16, 15, 20, 0, time.UTC),
	))

	fmt.Printf("Passage on easter costed %d\n", tollCalculator.GetTollFee(
		Car{},
		time.Date(2023, 4, 9, 14, 15, 0, 0, time.UTC),
		time.Date(2023, 4, 9, 15, 15, 10, 0, time.UTC),
		time.Date(2023, 4, 9, 16, 15, 20, 0, time.UTC),
	))
}

func setUpRate(start, end string, price int) toll.Rate {
	rate, err := toll.NewRate(start, end, price)
	if err != nil {
		panic(err)
	}

	return rate
}
