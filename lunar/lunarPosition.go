package lunar

import (
	"time"

	c "../common"
	o "../orbit"
	"../solar"
	t "../time"
)

const (
	lunarMeanLongitudeAtEpoch          = 318.351648 // degrees
	lunarMeanLongitudeOfPerigeeAtEpoch = 36.340410  // degrees
	lunarMeanLongitudeOfNodeAtEpoch    = 318.510107 // degrees
	lunarOrbitInclination              = 5.145396   // degrees
)

func Position(date time.Time) LunarPosition {
	D := t.EpochTime(date)
	Ms := solar.MeanAnomaly(date)
	Ls := solar.GeocentricEclipticLongitude(date)
	l := lunarLongitude(D)
	Mm := lunarAnomaly(l, D)
	N := lunarLongitudeOfNode(D)
	Ev := evectionCorrection(l, Ls, Mm)
	Ae := annualEquation(Ms)
	A3 := thirdCorrection(Ms)
	MMm := lunarCorrectedAnomaly(Mm, Ev, Ae, A3)
	Ec := centreEquationCorrection(MMm)
	A4 := fourthCorrection(MMm)
	ll := lunarCorrectedLongitude(l, Ev, Ec, Ae, A4)
	V := lunarVariation(ll, Ls)
	lll := lunarOrbitalLongitude(ll, V)
	NN := lunarCorrectedLongitudeOfNode(N, Ms)
	y := yCord(lll, NN)
	x := xCord(lll, NN)
	ec := o.Ecliptic{
		Latitude:  lunarEclipticLatitude(lll, NN),
		Longitude: lunarEclipticLongitude(x, y, NN),
	}

	debug := LunarPositionDebug{
		date: date,
		D:    D,
		Ms:   Ms,
		Ls:   Ls,
		l:    l,
		Mm:   Mm,
		N:    N,
		Ev:   Ev,
		Ae:   Ae,
		A3:   A3,
		MMm:  MMm,
		Ec:   Ec,
		A4:   A4,
		ll:   ll,
		V:    V,
		lll:  lll,
		NN:   NN,
		y:    y,
		x:    x,
	}

	return LunarPosition{Ecliptic: ec, Debug: debug}
}

func lunarLongitude(D float64) float64 {
	l := 13.1763966*D + lunarMeanLongitudeAtEpoch
	return c.NormaliseAngle(l)
}

func lunarAnomaly(l float64, D float64) float64 {
	M := l - 0.111404*D - lunarMeanLongitudeOfPerigeeAtEpoch
	return c.NormaliseAngle(M)
}

func lunarLongitudeOfNode(D float64) float64 {
	N := lunarMeanLongitudeOfNodeAtEpoch - (0.0529539 * D)
	return c.AdjustTo360(N)
}

func lunarCorrectedLongitudeOfNode(N float64, Ms float64) float64 {
	NN := N - (0.16 * c.Sind(Ms))
	return NN
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

func yCord(lll float64, NN float64) float64 {
	y := c.Sind(lll-NN) * c.Cosd(lunarOrbitInclination)
	return y
}

func xCord(lll float64, NN float64) float64 {
	x := c.Cosd(lll - NN)
	return x
}

func lunarEclipticLongitude(x float64, y float64, NN float64) float64 {
	Lm := c.Atan2d(y, x) + NN
	return Lm
}

func lunarEclipticLatitude(lll float64, NN float64) float64 {
	Bm := c.Asind(c.Sind(lll-NN) * c.Sind(lunarOrbitInclination))
	return Bm
}
