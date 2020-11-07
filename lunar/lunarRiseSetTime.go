package lunar

import (
	"fmt"
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
	GST1r               t.GST
	GST1s               t.GST
	GST2r               t.GST
	GST2s               t.GST
	GSTr                t.GST
	GSTs                t.GST
	dd                  o.Equatorial
	pi                  float64
	th                  float64
	x                   float64
	phi                 float64
	y                   float64
	dt                  float64
	GSTra               t.GST
	GSTsa               t.GST
	UTr                 time.Time
	UTs                 time.Time
}

func RiseTime(observer c.Observer, date time.Time) LunarRiseSetTime {
	midnightPosition := Position(time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC))
	middayPosition := Position(time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, time.UTC))

	midnightRiseSetTime := midnightPosition.Ecliptic.ToEquatorial(date).GetRiseSetTime(observer)
	middayRiseSetTime := middayPosition.Ecliptic.ToEquatorial(date).GetRiseSetTime(observer)

	UTdate := date.In(time.UTC)
	T00 := t.UTToGst(time.Date(UTdate.Year(), UTdate.Month(), UTdate.Day(), 0, 0, 0, 0, time.UTC))
	T000 := getT000(T00, observer)

	GST1r := calculateAdjustedGST(midnightRiseSetTime.LSTr, observer, T000)
	GST1s := calculateAdjustedGST(midnightRiseSetTime.LSTs, observer, T000)
	GST2r := calculateAdjustedGST(middayRiseSetTime.LSTr, observer, T000)
	GST2s := calculateAdjustedGST(middayRiseSetTime.LSTs, observer, T000)

	GSTr := interpolateGST(GST1r, GST2r, T00)
	GSTs := interpolateGST(GST1s, GST2s, T00)

	dd := meanPosition(midnightPosition, middayPosition, date)
	pi := horizontalParallax(midnightPosition)
	th := apparentAngularDiameter(midnightPosition)
	x := (-1.0 * pi) + (th / 2.0) + atmosphericRefraction
	phi := o.AngleAtHorizon(observer, dd)
	y := o.Y(x, phi)
	dt := o.RiseSetTimeShiftSeconds(y, dd) / 3600.0

	GSTra := t.GST(GSTr.Value() - dt)
	GSTsa := t.GST(GSTs.Value() + dt)

	UTr := GSTra.ToUT(date)
	UTs := GSTsa.ToUT(date)

	debug := LunarRiseSetTimeDebug{
		date:                date,
		observer:            observer,
		midnightPosition:    midnightPosition,
		middayPosition:      middayPosition,
		midnightRiseSetTime: midnightRiseSetTime,
		middayRiseSetTime:   middayRiseSetTime,
		T00:                 T00,
		T000:                T000,
		GST1r:               GST1r,
		GST1s:               GST1s,
		GST2r:               GST2r,
		GST2s:               GST2s,
		GSTr:                GSTr,
		GSTs:                GSTs,
		dd:                  dd,
		pi:                  pi,
		th:                  th,
		x:                   x,
		phi:                 phi,
		y:                   y,
		dt:                  dt,
		GSTra:               GSTra,
		GSTsa:               GSTsa,
		UTr:                 UTr,
		UTs:                 UTs,
	}

	return LunarRiseSetTime{
		Rise:  UTr.In(observer.Location),
		Set:   UTs.In(observer.Location),
		Debug: debug,
	}
}

func getT000(T00 t.GST, observer c.Observer) t.GST {
	T000 := T00.Value() - ((observer.Longitude / 15.0) * 1.002738)
	if T000 < 0 {
		T000 += 24.0
	}
	return t.GST(T000)
}

func calculateAdjustedGST(LST t.LST, observer c.Observer, T000 t.GST) t.GST {
	GSTvalue := LST.ToGst(observer).Value()
	fmt.Printf("GST: %f\n", GSTvalue)
	if GSTvalue < T000.Value() {
		GSTvalue += 24.0
	}
	return t.GST(GSTvalue)
}

func interpolateGST(GST1 t.GST, GST2 t.GST, T00 t.GST) t.GST {
	GST := ((12.03 * GST1.Value()) - (T00.Value() * (GST2.Value() - GST1.Value()))) / (12.03 + GST1.Value() - GST2.Value())
	return t.GST(GST)
}

func meanPosition(p1 LunarPosition, p2 LunarPosition, date time.Time) o.Equatorial {
	return o.MeanEquatorialPosition(p1.Ecliptic.ToEquatorial(date), p2.Ecliptic.ToEquatorial(date))
}
