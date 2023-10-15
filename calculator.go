package main

import (
	"time"
	"toll-fee-calculator/toll"
)

type Calculator struct {
	config      []toll.Config
	maxPrice    int
	gracePeriod time.Duration
}

// GetTollFee Calculate the total toll fee for one day given a vehicle and the timestamps
// of all the passages during that day.
func (c *Calculator) GetTollFee(vehicle Vehicle, dates ...time.Time) int {
	if isTollFreeVehicle(vehicle) || len(dates) == 0 {
		return 0
	}
	startOfDay := dates[0].Truncate(time.Hour * 24)

	tolls := make(map[time.Duration]int, len(dates))
	var keys []time.Duration
	for _, date := range dates {
		keys = append(keys, date.Sub(startOfDay))
		tolls[date.Sub(startOfDay)] = c.tollFeePassage(date)
	}

	totalFee := 0
	tempHighFee := 0
	offsetCursor, _ := time.ParseDuration("0m")
	for _, passageOffset := range keys {
		if passageOffset-offsetCursor <= c.gracePeriod {
			if tolls[passageOffset] > tempHighFee {
				tempHighFee = tolls[passageOffset]
			}
		} else {
			totalFee += tempHighFee
			tempHighFee = tolls[passageOffset]
			offsetCursor = passageOffset
		}
	}
	totalFee += tempHighFee

	if totalFee > c.maxPrice {
		return c.maxPrice
	}

	return totalFee
}

func (c *Calculator) tollFeePassage(date time.Time) int {
	for _, tollConfig := range c.config {
		if tollConfig.AppliesTo(date) {
			return tollConfig.GetPrice(date)
		}
	}

	return 0 // Return 0 if nothing matches
}

func NewCalculator(config []toll.Config) Calculator {
	gracePeriod, _ := time.ParseDuration("60m")
	return Calculator{config: config, maxPrice: 60, gracePeriod: gracePeriod}
}
