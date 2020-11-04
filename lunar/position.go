package lunar

import (
	c "../common"
)

type LunarPosition struct {
	Ecliptic c.Ecliptic
	Debug    LunarPositionDebug // Calculation debug values
}

type LunarPositionDebug struct {
	D   float64 // Epoch Time; days
	Ms  float64 // Solar Mean Anomaly
	Ls  float64 // Solar Geocentric Ecliptic Longitude
	l   float64 // Lunar Orbital Longitude
	Mm  float64 // Lunar Mean Anomaly
	N   float64 // Lunar Longitude of Node; degrees
	Ev  float64 // Evection Correction
	Ae  float64 // Lunar Annual Equation
	A3  float64 // Lunar Third Correction to Mean Anomaly
	MMm float64
	Ec  float64 // Lunar Correction to Equation of Centre
	A4  float64 // Lunar Fourth Correction to Mean Anomaly
	ll  float64 // Lunar Corrected Longitude
	V   float64 // Lunar Variation
	lll float64 // Lunar True Longitude
	NN  float64
	y   float64
	x   float64
}
