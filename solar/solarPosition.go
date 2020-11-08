package solar

import (
	"math"
	"time"

	c "github.com/rtovey/astro/common"
	o "github.com/rtovey/astro/orbit"
)

const (
	solarYearDurationDays           = 365.242191
	solarEclipticLongitudeAtEpoch   = 279.403303
	solarEclipticLongitudeOfPerigee = 282.768422
	solarOrbitEccentricity          = 0.016713
)

type SolarPosition struct {
	Ecliptic o.Ecliptic
	Debug    SolarPositionDebug // Calculation debug values
}

type SolarPositionDebug struct {
	D float64 // days
	N float64 // degrees
	M float64 // degrees
	E float64 // radians
	v float64 // degrees
}

func Position(date time.Time) SolarPosition {
	D := daysSinceEpoch(date)
	N := northPointOFHorizon(D)
	M := meanAnomaly(N)
	E := c.SolveKeplersEquation(c.DtoR(M), solarOrbitEccentricity, math.Pow10(-6))
	v := trueAnomaly(E)
	l := v + solarEclipticLongitudeOfPerigee
	l = c.NormaliseAngle(l)

	debug := SolarPositionDebug{
		D: D,
		N: N,
		M: M,
		E: E,
		v: v,
	}

	return SolarPosition{
		Ecliptic: o.Ecliptic{
			Latitude:  0,
			Longitude: l,
		},
		Debug: debug,
	}
}

func daysSinceEpoch(date time.Time) float64 {
	epoch := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	return (date.Sub(epoch).Hours() + 24.0) / 24.0
}

func northPointOFHorizon(D float64) float64 {
	return c.NormaliseAngle((360.0 / solarYearDurationDays) * D)
}

func meanAnomaly(N float64) float64 {
	M := N + solarEclipticLongitudeAtEpoch - solarEclipticLongitudeOfPerigee
	if M < 0 {
		M += 360.0
	}
	return M
}

func trueAnomaly(E float64) float64 {
	x := math.Pow((1+solarOrbitEccentricity)/(1-solarOrbitEccentricity), 0.5) * math.Tan(E/2.0)
	return c.RtoD(2 * math.Atan(x))
}

/*func equationOfCentreCorrection(date time.Time) float64 {
	M := MeanAnomaly(date)
	return (360.0 / math.Pi) * solarOrbitEccentricity * c.Sind(M)
}

func GeocentricEclipticLongitude(date time.Time) float64 {
	N := northPointOFHorizon(date)
	Ec := equationOfCentreCorrection(date)

	L := N + Ec + solarEclipticLongitudeAtEpoch
	if L > 360 {
		L = L - 360.0
	}
	return L
}*/
