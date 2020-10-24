package solar

import (
	"fmt"
	"math"
	"time"

	c "../common"
	astroTime "../time"
)

const (
	solarYearDurationDays           = 365.242191
	solarEclipticLongitudeAtEpoch   = 279.403303
	solarEclipticLongitudeOfPerigee = 282.768422
	solarOrbitEccentricity          = 0.016713
)

func northPointOFHorizon(date time.Time) float64 {
	D := astroTime.DSinceEpoch(date)
	fmt.Printf("D = %f\n", D)
	N := (360.0 / solarYearDurationDays) * D
	return c.AdjustTo360(N)
}

func MeanAnomaly(date time.Time) float64 {
	N := northPointOFHorizon(date)
	fmt.Printf("N = %f\n", N)
	M := N + solarEclipticLongitudeAtEpoch - solarEclipticLongitudeOfPerigee
	if M < 0 {
		M = M + 360.0
	}
	return M
}

func equationOfCentreCorrection(date time.Time) float64 {
	M := MeanAnomaly(date)
	return (360.0 / math.Pi) * solarOrbitEccentricity * c.Sind(M)
}

func GeocentricEclipticLongitude(date time.Time) float64 {
	N := northPointOFHorizon(date)
	Ec := equationOfCentreCorrection(date)
	fmt.Printf("Ec = %f\n", Ec)

	L := N + Ec + solarEclipticLongitudeAtEpoch
	fmt.Printf("L = %f\n", L)
	if L > 360 {
		L = L - 360.0
	}
	return L
}
