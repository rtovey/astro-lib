package time

import (
	"time"
)

const (
	// Solar
	//epochYear        = 2010
	etOffsetSeconds = 66.07
	epochYear       = 1990
	epochMonth      = time.January
	epochDay        = 0
	epochHour       = 0
	epochMinute     = 0
	epochSecond     = 0
	epochNSecond    = 0
)

func daysSinceEpoch(date time.Time) int {
	daysToCurrentYear := 0

	if date.Year() > epochYear {
		for i := epochYear; i != date.Year(); i++ {
			if isLeapYear(i) {
				daysToCurrentYear += 366
			} else {
				daysToCurrentYear += 365
			}
		}
	}
	if date.Year() < epochYear {
		for i := date.Year(); i != epochYear; i++ {
			if isLeapYear(i) {
				daysToCurrentYear -= 366
			} else {
				daysToCurrentYear -= 365
			}
		}
	}
	return date.YearDay() + daysToCurrentYear
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func EpochTime(date time.Time) float64 {
	utcDate := date.UTC()
	epochT := (float64(utcDate.Hour()) / 24.0) +
		(float64(utcDate.Minute()) / 1440.0) +
		((float64(utcDate.Second()) + etOffsetSeconds) / 86400.0) +
		float64(daysSinceEpoch(utcDate))
	return epochT
}

func DSinceEpoch(date time.Time) float64 {
	epoch := time.Date(epochYear, epochMonth, epochDay, epochHour, epochMinute, epochSecond, epochNSecond, time.UTC)
	return date.Sub(epoch).Hours() / 24
}
