package common

import "time"

type Equatorial struct {
	RightAscension float64
	Declination    float64
	Debug          EquatorialDebug
}

type EquatorialDebug struct {
	ecliptic Ecliptic
	e        float64
	d        float64
	x        float64
	y        float64
	a        float64
}

func (ec Ecliptic) ToEquatorial(date time.Time) Equatorial {
	e := MeanObliquityOfEcliptic(date)
	d := Asind((Sind(ec.Latitude) * Cosd(e)) + (Cosd(ec.Latitude) * Sind(e) * Sind(ec.Longitude)))
	y := (Sind(ec.Longitude) * Cosd(e)) - (Tand(ec.Latitude) * Sind(e))
	x := Cosd(ec.Longitude)
	a := Atan2d(y, x) / 15.0

	debug := EquatorialDebug{
		ecliptic: ec,
		e:        e,
		d:        d,
		x:        x,
		y:        y,
		a:        a,
	}

	return Equatorial{
		RightAscension: a,
		Declination:    d,
		Debug:          debug,
	}
}

func MeanObliquityOfEcliptic(date time.Time) float64 {
	// TODO
	return 23.441884
}
