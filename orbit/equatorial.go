package orbit

import (
	"time"

	c "../common"
)

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
	d := c.Asind((c.Sind(ec.Latitude) * c.Cosd(e)) + (c.Cosd(ec.Latitude) * c.Sind(e) * c.Sind(ec.Longitude)))
	y := (c.Sind(ec.Longitude) * c.Cosd(e)) - (c.Tand(ec.Latitude) * c.Sind(e))
	x := c.Cosd(ec.Longitude)
	a := c.Atan2d(y, x) / 15.0

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
	//return 23.441884
	return 23.43658135
}
