package solar

import (
	"time"

	c "github.com/rtovey/astro/common"
	o "github.com/rtovey/astro/orbit"
	t "github.com/rtovey/astro/time"
)

type SolarRiseSetTime struct {
	Rise  time.Time
	Set   time.Time
	Debug SolarRiseSetTimeDebug
}

type SolarRiseSetTimeDebug struct {
	date                 time.Time
	observer             c.Observer
	midnightPosition     SolarPosition
	plus24HrsPosition    SolarPosition
	midnightRiseSetTime  o.RiseSetTime
	plus24HrsRiseSetTime o.RiseSetTime
	T00                  t.GST
	T000                 t.GST
	GST1r                t.GST
	GST1s                t.GST
	GST2r                t.GST
	GST2s                t.GST
	GSTr                 t.GST
	GSTs                 t.GST
	dd                   o.Equatorial
	pi                   float64
	th                   float64
	x                    float64
	phi                  float64
	y                    float64
	dt                   float64
	GSTra                t.GST
	GSTsa                t.GST
	UTr                  time.Time
	UTs                  time.Time
}

func RiseSetTime(observer c.Observer, date time.Time) SolarRiseSetTime {
	midnightUTDate := date.In(time.UTC)
	plus24HrsUTDate := midnightUTDate.Add(time.Hour * 24)

	midnightPosition := Position(time.Date(midnightUTDate.Year(), midnightUTDate.Month(), midnightUTDate.Day(), 0, 0, 0, 0, time.UTC))
	plus24HrsPosition := Position(time.Date(plus24HrsUTDate.Year(), plus24HrsUTDate.Month(), plus24HrsUTDate.Day(), 0, 0, 0, 0, time.UTC))

	T00 := t.UTToGst(time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC))
	T000 := t.GetT000(T00, observer)

	midnightRiseSetTime := midnightPosition.Ecliptic.ToEquatorial(date).GetRiseSetTime(observer)
	plus24HrsRiseSetTime := plus24HrsPosition.Ecliptic.ToEquatorial(date).GetRiseSetTime(observer)

	GST1r := midnightRiseSetTime.LSTr.ToGst(observer)
	GST1s := midnightRiseSetTime.LSTs.ToGst(observer)
	GST2r := plus24HrsRiseSetTime.LSTr.ToGst(observer)
	GST2s := plus24HrsRiseSetTime.LSTs.ToGst(observer)

	if GST1r.Value() > GST2r.Value() {
		GST2r = t.GST(GST2r.Value() + 24.0)
	}
	if GST1s.Value() > GST2s.Value() {
		GST2s = t.GST(GST2s.Value() + 24.0)
	}
	if GST1r.Value() < T000.Value() {
		GST1r = t.GST(GST1r.Value() + 24.0)
		GST2r = t.GST(GST2r.Value() + 24.0)
	}
	if GST1s.Value() < T000.Value() {
		GST1s = t.GST(GST1s.Value() + 24.0)
		GST2s = t.GST(GST2s.Value() + 24.0)
	}

	GSTr := interpolateGST(GST1r, GST2r, T00)
	GSTs := interpolateGST(GST1s, GST2s, T00)

	dd := meanPosition(midnightPosition, plus24HrsPosition, date)
	pi := horizontalParallax
	th := apparentAngularDiameter(midnightPosition)
	x := (-1.0 * pi) + (th / 2.0) + atmosphericRefraction
	phi := o.AngleAtHorizon(observer, dd)
	y := o.Y(x, phi)
	dt := o.RiseSetTimeShiftSeconds(y, dd) / 3600.0

	GSTra := t.GST(GSTr.Value() - dt)
	GSTsa := t.GST(GSTs.Value() + dt)

	UTr := GSTra.ToUT(date)
	UTs := GSTsa.ToUT(date)

	debug := SolarRiseSetTimeDebug{
		date:                 date,
		observer:             observer,
		midnightPosition:     midnightPosition,
		plus24HrsPosition:    plus24HrsPosition,
		midnightRiseSetTime:  midnightRiseSetTime,
		plus24HrsRiseSetTime: plus24HrsRiseSetTime,
		T00:                  T00,
		T000:                 T000,
		GST1r:                GST1r,
		GST1s:                GST1s,
		GST2r:                GST2r,
		GST2s:                GST2s,
		GSTr:                 GSTr,
		GSTs:                 GSTs,
		dd:                   dd,
		pi:                   pi,
		th:                   th,
		x:                    x,
		phi:                  phi,
		y:                    y,
		dt:                   dt,
		GSTra:                GSTra,
		GSTsa:                GSTsa,
		UTr:                  UTr,
		UTs:                  UTs,
	}

	return SolarRiseSetTime{
		Rise:  UTr.In(observer.Location),
		Set:   UTs.In(observer.Location),
		Debug: debug,
	}
}

func interpolateGST(GST1 t.GST, GST2 t.GST, T00 t.GST) t.GST {
	GST := ((24.07 * GST1.Value()) - (T00.Value() * (GST2.Value() - GST1.Value()))) / (24.07 + GST1.Value() - GST2.Value())
	return t.GST(GST)
}

func meanPosition(p1 SolarPosition, p2 SolarPosition, date time.Time) o.Equatorial {
	return o.MeanEquatorialPosition(p1.Ecliptic.ToEquatorial(date), p2.Ecliptic.ToEquatorial(date))
}
