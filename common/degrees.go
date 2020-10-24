package common

import (
	"math"
)

func Sind(degrees float64) float64 {
	return math.Sin(DtoR(degrees))
}

func Cosd(degrees float64) float64 {
	return math.Cos(DtoR(degrees))
}

func DtoR(degrees float64) float64 {
	return (math.Pi / 180) * degrees
}

func AdjustTo360(degrees float64) float64 {
	for degrees < 0.0 {
		degrees = degrees + 360.0
	}
	for degrees > 360.0 {
		degrees = degrees - 360.0
	}
	return degrees
}
