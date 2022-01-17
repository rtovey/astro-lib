package time

import (
	"math"
	"time"

	c "github.com/rtovey/astro-lib/common"
)

type LST float64
type GST float64

func (lst LST) Value() float64 {
	return float64(lst)
}

func (gst GST) Value() float64 {
	return float64(gst)
}

func (lst LST) ToGst(observer c.Observer) GST {
	gst := lst.Value()
	timeDiff := observer.Longitude / 15.0
	gst -= timeDiff

	if gst > 24.0 {
		gst -= 24.0
	}
	if gst < 0 {
		gst += 24.0
	}
	return GST(gst)
}

func UTToGst(date time.Time) GST {
	UTdate := date.In(time.UTC)
	JD := ToJulianDate(time.Date(UTdate.Year(), UTdate.Month(), UTdate.Day(), 0, 0, 0, 0, time.UTC)).Value()

	S := JD - 2451545.0
	T := S / 36525.0
	T0 := 6.697374558 + (2400.051336 * T) + (0.000025862 * math.Pow(T, 2))
	T0 = adjustTo24Hours(T0)
	UT := toDecimalHours(date)

	T0 += UT * 1.002737909
	T0 = adjustTo24Hours(T0)
	return GST(T0)
}

func (gst GST) ToUT(date time.Time) time.Time {
	UTdate := date.In(time.UTC)
	JD := ToJulianDate(time.Date(UTdate.Year(), UTdate.Month(), UTdate.Day(), 0, 0, 0, 0, time.UTC)).Value()

	S := JD - 2451545.0
	T := S / 36525.0
	T0 := 6.697374558 + (2400.051336 * T) + (0.000025862 * math.Pow(T, 2))
	T0 = adjustTo24Hours(T0)
	T0 = gst.Value() - T0
	T0 = adjustTo24Hours(T0)

	UT := Hours(0.9972695663 * T0)
	return time.Date(UTdate.Year(), UTdate.Month(), UTdate.Day(), UT.Hours(), UT.Minutes(), UT.Seconds(), 0, time.UTC)
}

func GetT000(T00 GST, observer c.Observer) GST {
	T000 := T00.Value() - ((observer.Longitude / 15.0) * 1.002738)
	if T000 < 0 {
		T000 += 24.0
	}
	return GST(T000)
}

func adjustTo24Hours(hours float64) float64 {
	for hours < 0.0 {
		hours = hours + 24.0
	}
	for hours > 24.0 {
		hours = hours - 24.0
	}
	return hours
}

func toDecimalHours(date time.Time) float64 {
	UTdate := date.In(time.UTC)
	return float64(UTdate.Hour()) + (float64(UTdate.Minute()) / 60.0) + (float64(UTdate.Second()) / 3600.0)
}
