package orbit

import (
	c "../common"
	t "../time"
)

type RiseSetTime struct {
	LSTr  t.LST
	LSTs  t.LST
	Debug RiseSetDebug
}

type RiseSetDebug struct {
	Position Equatorial
	Latitude float64 // observer's latitude; degrees
	Ar       float64 // degrees
	As       float64 // degrees
	H        float64 // hours
}

func (object Equatorial) GetRiseSetTime(observer c.Observer) RiseSetTime {
	latitude := observer.Latitude
	Ar := c.Acosd(c.Sind(object.Declination) / c.Cosd(latitude))
	As := 360.0 - Ar
	H := (1 / 15.0) * c.Acosd(-1*c.Tand(latitude)*c.Tand(object.Declination))
	LSTr := t.LST(localSiderialRiseTime(object, H))
	LSTs := t.LST(localSiderialSetTime(object, H))

	debug := RiseSetDebug{
		Position: object,
		Latitude: latitude,
		Ar:       Ar,
		As:       As,
		H:        H,
	}

	return RiseSetTime{
		LSTr:  LSTr,
		LSTs:  LSTs,
		Debug: debug,
	}
}

func localSiderialRiseTime(position Equatorial, H float64) float64 {
	LSTr := 24.0 + position.RightAscension - H
	if LSTr > 24 {
		LSTr = LSTr - 24.0
	}
	return LSTr
}

func localSiderialSetTime(position Equatorial, H float64) float64 {
	LSTs := position.RightAscension + H
	if LSTs > 24 {
		LSTs = LSTs - 24.0
	}
	return LSTs
}
