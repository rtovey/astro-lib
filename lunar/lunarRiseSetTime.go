package lunar

import (
	"time"

	c "../common"
	o "../orbit"
)

type LunarRiseSetTime struct {
	Rise  time.Time
	Set   time.Time
	Debug LunarRiseSetTimeDebug
}

type LunarRiseSetTimeDebug struct {
	date                time.Time
	observer            c.Observer
	midnightPosition    LunarPosition
	middayPosition      LunarPosition
	midnightRiseSetTime o.RiseSetTime
	middayRiseSetTime   o.RiseSetTime
}

func RiseTime(observer c.Observer, date time.Time) LunarRiseSetTime {
	midnightPosition := Position(time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC))
	middayPosition := Position(time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, time.UTC))

	midnightRiseSetTime := midnightPosition.Ecliptic.ToEquatorial(date).GetRiseSetTime(observer)
	middayRiseSetTime := middayPosition.Ecliptic.ToEquatorial(date).GetRiseSetTime(observer)

	debug := LunarRiseSetTimeDebug{
		date:                date,
		observer:            observer,
		midnightPosition:    midnightPosition,
		middayPosition:      middayPosition,
		midnightRiseSetTime: midnightRiseSetTime,
		middayRiseSetTime:   middayRiseSetTime,
	}

	return LunarRiseSetTime{
		Rise:  time.Now(),
		Set:   time.Now(),
		Debug: debug,
	}
}
