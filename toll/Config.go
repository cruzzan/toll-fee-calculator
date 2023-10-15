package toll

import (
	"time"
	"toll-fee-calculator/dayMatcher"
)

type Config struct {
	rates   []Rate
	matcher dayMatcher.DayMatcher
}

// AppliesTo checks if this config applies to the given instant, using the configured DayMatcher
func (b Config) AppliesTo(instant time.Time) bool {
	return b.matcher.Matches(instant)
}

// GetPrice looks up the price of the instant in the configured rates.
// Defaults to 0 if no rates are configured
func (b Config) GetPrice(instant time.Time) int {
	for _, rate := range b.rates {
		if rate.Matches(instant) {
			return rate.GetPrice()
		}
	}
	return 0
}

func NewConfig(matcher dayMatcher.DayMatcher, rates ...Rate) Config {
	return Config{matcher: matcher, rates: rates}
}
