package main

import "time"

func get2023Holidays() []time.Time {
	return []time.Time{
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),   // New years day
		time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC),   // Epiphany
		time.Date(2023, 4, 7, 0, 0, 0, 0, time.UTC),   // Good friday
		time.Date(2023, 4, 9, 0, 0, 0, 0, time.UTC),   // Easter
		time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),  // Easter monday
		time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),   // May day
		time.Date(2023, 5, 18, 0, 0, 0, 0, time.UTC),  // Ascension day
		time.Date(2023, 5, 28, 0, 0, 0, 0, time.UTC),  // Pentecost day
		time.Date(2023, 6, 6, 0, 0, 0, 0, time.UTC),   // National day
		time.Date(2023, 6, 24, 0, 0, 0, 0, time.UTC),  // Midsummer day
		time.Date(2023, 11, 4, 0, 0, 0, 0, time.UTC),  // All hallows day
		time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC), // Christmas day
		time.Date(2023, 12, 26, 0, 0, 0, 0, time.UTC), // Boxing day
	}
}

func get2023DayBeforeHolidays() []time.Time {
	return []time.Time{
		time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),   // Epiphany eve
		time.Date(2023, 4, 6, 0, 0, 0, 0, time.UTC),   // Maundy thursday
		time.Date(2023, 4, 8, 0, 0, 0, 0, time.UTC),   // Easter eve
		time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC),  // Walpurgis night
		time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),  // Day before ascension day
		time.Date(2023, 5, 27, 0, 0, 0, 0, time.UTC),  // Pentecost eve
		time.Date(2023, 6, 5, 0, 0, 0, 0, time.UTC),   // Day before National day
		time.Date(2023, 6, 23, 0, 0, 0, 0, time.UTC),  // Midsummer eve
		time.Date(2023, 11, 3, 0, 0, 0, 0, time.UTC),  // All hallows eve
		time.Date(2023, 12, 24, 0, 0, 0, 0, time.UTC), // Christmas Eve
		time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC), // New years eve
	}
}
