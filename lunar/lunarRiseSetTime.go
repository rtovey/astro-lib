package lunar

import (
	"time"

	c "../common"
	o "../orbit"
	t "../time"
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
	T00                 t.GST
	T000                t.GST
}

func RiseTime(observer c.Observer, date time.Time) LunarRiseSetTime {
	midnightPosition := Position(time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC))
	middayPosition := Position(time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, time.UTC))

	midnightRiseSetTime := midnightPosition.Ecliptic.ToEquatorial(date).GetRiseSetTime(observer)
	middayRiseSetTime := middayPosition.Ecliptic.ToEquatorial(date).GetRiseSetTime(observer)

	UTdate := date.In(time.UTC)
	T00 := t.UTToGst(time.Date(UTdate.Year(), UTdate.Month(), UTdate.Day(), 0, 0, 0, 0, time.UTC))
	T000 := T00.Value() - ((observer.Longitude / 15.0) * 1.002738)
	if T000 < 0 {
		T000 += 24.0
	}

	debug := LunarRiseSetTimeDebug{
		date:                date,
		observer:            observer,
		midnightPosition:    midnightPosition,
		middayPosition:      middayPosition,
		midnightRiseSetTime: midnightRiseSetTime,
		middayRiseSetTime:   middayRiseSetTime,
		T00:                 T00,
		T000:                t.GST(T000),
	}

	return LunarRiseSetTime{
		Rise:  time.Now(),
		Set:   time.Now(),
		Debug: debug,
	}
}
