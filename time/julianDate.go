package time

import (
	"math"
	"time"
)

type JulianDate float64

const (
	julianDateOffset = 1720994.5
)

func ToJulianDate(t time.Time) JulianDate {
	gregorianCalendarStart := time.Date(1582, time.October, 14, 23, 59, 59, 999, time.UTC)

	y := float64(t.Year())
	m := float64(t.Month())
	if m == 1 || m == 2 {
		y--
		m += 12
	}
	d := d(t)

	A, _ := math.Modf(y / 100.0)

	B := 0.0
	if t.After(gregorianCalendarStart) {
		A4, _ := math.Modf(A / 4.0)
		B = 2.0 - A + A4
	}

	C := 0.0
	if y < 0 {
		C, _ = math.Modf((365.25 * y) - 0.75)
	} else {
		C, _ = math.Modf(365.25 * y)
	}

	D, _ := math.Modf(30.6001 * (m + 1))

	return JulianDate(B + C + D + d + julianDateOffset)
}

func d(t time.Time) float64 {
	return float64(t.Day()) +
		(float64(t.Hour()) / 24.0) +
		(float64(t.Minute() / 3600.0)) +
		(float64(t.Second()) / 86400.0)
}

func (JD JulianDate) Value() float64 {
	return float64(JD)
}
