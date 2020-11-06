package time

import (
	c "../common"
)

type RiseSetTime struct {
	LSTr  LocalSiderealTime
	LSTs  LocalSiderealTime
	Debug RiseSetTimeDebug
}

type RiseSetTimeDebug struct {
	Position c.Equatorial
	Latitude float64 // observer's latitude; degrees
	Ar       float64 // degrees
	As       float64 // degrees
	H        float64 // hours
}

func (p c.Equatorial) RiseSetTime(observer c.Observer) RiseSetTime {
	latitude := observer.Latitude
	Ar := Acosd(Sind(p.Declination) / Cosd(latitude))
	As := 360.0 - Ar
	H := (1 / 15.0) * Acosd(-1*Tand(latitude)*Tand(p.Declination))
	LSTr := t.LocalSiderealTime{
		Time: localSiderialRiseTime(p, H),
	}
	LSTs := t.LocalSiderealTime{
		Time: localSiderialSetTime(p, H),
	}

	debug := RiseSetTimeDebug{
		Position: p,
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

func localSiderialRiseTime(position c.Equatorial, H float64) float64 {
	LSTr := 24.0 + position.RightAscension - H
	if LSTr > 24 {
		LSTr = LSTr - 24.0
	}
	return LSTr
}

func localSiderialSetTime(position c.Equatorial, H float64) float64 {
	LSTs := position.RightAscension + H
	if LSTs > 24 {
		LSTs = LSTs - 24.0
	}
	return LSTs
}
