package solar

import (
	"time"

	c "github.com/rtovey/astro/common"
)

type SolarRiseSetTime struct {
	Rise  time.Time
	Set   time.Time
	Debug SolarRiseSetTimeDebug
}

type SolarRiseSetTimeDebug struct {
	date              time.Time
	observer          c.Observer
	midnightPosition  SolarPosition
	plus24HrsPosition SolarPosition
}

func RiseSetTime(observer c.Observer, date time.Time) SolarRiseSetTime {
	midnightUTDate := date.In(time.UTC)
	plus24HrsUTDate := midnightUTDate.Add(time.Hour * 24)

	midnightPosition := Position(time.Date(midnightUTDate.Year(), midnightUTDate.Month(), midnightUTDate.Day(), 0, 0, 0, 0, time.UTC))
	plus24HrsPosition := Position(time.Date(plus24HrsUTDate.Year(), plus24HrsUTDate.Month(), plus24HrsUTDate.Day(), 0, 0, 0, 0, time.UTC))

	debug := SolarRiseSetTimeDebug{
		date:              date,
		observer:          observer,
		midnightPosition:  midnightPosition,
		plus24HrsPosition: plus24HrsPosition,
	}

	return SolarRiseSetTime{
		Debug: debug,
	}
}
