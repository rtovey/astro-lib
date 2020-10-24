package lunar

import (
	"fmt"
	"time"

	c "../common"
	solar "../solar"
	t "../time"
)

const (
	lunarMeanLongitudeAtEpoch          = 318.351648 // degrees
	lunarMeanLongitudeOfPerigeeAtEpoch = 36.340410  // degrees
	lunarMeanLongitudeOfNodeAtEpoch    = 318.510107 // degrees
)

func RiseTime(date time.Time) time.Time {
	Ls := solar.GeocentricEclipticLongitude(date)
	Ms := solar.MeanAnomaly(date)
	l := meanLongitude(date)
	Ml := meanAnomaly(date, l)
	N := northPointOfHorizon(date)
	Ev := evectionCorrection(l, Ml, Ls)

	fmt.Printf("Ls = %f\n", Ls)
	fmt.Printf("Ms = %f\n", Ms)
	fmt.Printf("l = %f\n", l)
	fmt.Printf("Ml = %f\n", Ml)
	fmt.Printf("N = %f\n", N)
	fmt.Printf("Ev = %f\n", Ev)
	return time.Now()
}

func meanLongitude(date time.Time) float64 {
	D := t.DSinceEpoch(date)
	l := 13.1763966*D + lunarMeanLongitudeAtEpoch
	return c.AdjustTo360(l)
}

func meanAnomaly(date time.Time, meanLongitude float64) float64 {
	D := t.DSinceEpoch(date)
	M := meanLongitude - (0.1114041 * D) - lunarMeanLongitudeOfPerigeeAtEpoch

	return c.AdjustTo360(M)
}

func northPointOfHorizon(date time.Time) float64 {
	D := t.DSinceEpoch(date)
	N := lunarMeanLongitudeOfNodeAtEpoch - (0.0529539 * D)
	return c.AdjustTo360(N)
}

func evectionCorrection(lunarMeanLongitude float64, lunarMeanAnomaly float64, solarGeocentricEclipticLongitude float64) float64 {
	return 1.2739 * c.Sind(2*(lunarMeanLongitude-solarGeocentricEclipticLongitude)-lunarMeanAnomaly)
}
