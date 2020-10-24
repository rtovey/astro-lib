package lunar

import (
	"math"
	"time"

	c "../common"
	t "../time"
)

const (
	// Solar
	SOLAR_ECLIPTIC_EPOCH_LONGITUDE   = 279.557208
	SOLAR_ECLIPTIC_PERIGEE_LONGITUDE = 283.112438
	SOLAR_ECCENTRICITY               = 0.016705
	solarYearDurationDays            = 365.242191

	// Lunar
	LUNAR_MEAN_EPOCH_LONGITUDE    = 91.929336
	LUNAR_PERIGEE_EPOCH_LONGITUDE = 130.143076
)

func LunarPhase(date time.Time) float64 {
	D := t.EpochTime(date)
	N := sunN(D)
	M := sunMeanAnomaly(N)
	L := sunLongitude(M, N)
	l := lunarLongitude(D)
	m := lunarAnomaly(l, D)
	Ev := evectionCorrection(l, L, m)
	Ae := annualEquation(M)
	A3 := thirdCorrection(M)
	mm := lunarCorrectedAnomaly(m, Ev, Ae, A3)
	Ec := centreEquationCorrection(mm)
	A4 := fourthCorrection(mm)
	ll := lunarCorrectedLongitude(l, Ev, Ec, Ae, A4)
	V := lunarVariation(ll, L)
	lll := lunarOrbitalLongitude(ll, V)
	F := lunarPhase(lll, L)
	return F
}

func normaliseAngle(angle float64) float64 {
	for angle < 0 || angle > 360 {
		var direction int
		if angle < 0 {
			direction = 1
		} else {
			direction = -1
		}
		angle += float64(direction * 360)
	}
	return angle
}

func sunN(epochTime float64) float64 {
	N := (360 / solarYearDurationDays) * epochTime
	return normaliseAngle(N)
}

func sunMeanAnomaly(N float64) float64 {
	meanAnomaly := N + SOLAR_ECLIPTIC_EPOCH_LONGITUDE - SOLAR_ECLIPTIC_PERIGEE_LONGITUDE
	if meanAnomaly < 0 {
		meanAnomaly += 360
	}
	return meanAnomaly
}

func sunLongitude(M float64, N float64) float64 {
	Ec := (360 / math.Pi) * SOLAR_ECCENTRICITY * c.Sind(M)
	l := N + Ec + SOLAR_ECLIPTIC_EPOCH_LONGITUDE

	if l > 360 {
		l -= 360
	}
	return l
}

func lunarLongitude(epochTime float64) float64 {
	l := 13.1763966*epochTime + LUNAR_MEAN_EPOCH_LONGITUDE
	return normaliseAngle(l)
}

func lunarAnomaly(l float64, D float64) float64 {
	M := l - 0.111404*D - LUNAR_PERIGEE_EPOCH_LONGITUDE
	return normaliseAngle(M)
}

func evectionCorrection(l float64, L float64, m float64) float64 {
	C := l - L
	return 1.2739 * c.Sind(2*C-m)
}

func annualEquation(M float64) float64 {
	return 0.1858 * c.Sind(M)
}

func thirdCorrection(M float64) float64 {
	return 0.37 * c.Sind(M)
}

func lunarCorrectedAnomaly(m float64, Ev float64, Ae float64, A3 float64) float64 {
	return m + Ev - Ae - A3
}

func centreEquationCorrection(mm float64) float64 {
	return 6.2886 * c.Sind(mm)
}

func fourthCorrection(mm float64) float64 {
	return 0.214 * c.Sind(2.0*mm)
}

func lunarCorrectedLongitude(l float64, Ev float64, Ec float64, Ae float64, A4 float64) float64 {
	return l + Ev + Ec - Ae + A4
}

func lunarVariation(ll float64, L float64) float64 {
	return 0.6583 * c.Sind(2.0*(ll-L))
}

func lunarOrbitalLongitude(ll float64, V float64) float64 {
	return ll + V
}

func lunarPhase(lll float64, L float64) float64 {
	return 0.5 * (1 - c.Cosd(lll-L))
}
