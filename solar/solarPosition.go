package solar

import (
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
	N := (360.0 / solarYearDurationDays) * D
	return c.AdjustTo360(N)
}

func MeanAnomaly(date time.Time) float64 {
	N := northPointOFHorizon(date)
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

	L := N + Ec + solarEclipticLongitudeAtEpoch
	if L > 360 {
		L = L - 360.0
	}
	return L
}
