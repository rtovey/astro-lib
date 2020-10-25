package common

import (
	"math"
)

func Sind(degrees float64) float64 {
	return math.Sin(dtoR(degrees))
}

func Cosd(degrees float64) float64 {
	return math.Cos(dtoR(degrees))
}

func dtoR(degrees float64) float64 {
	return (math.Pi / 180) * degrees
}

func NormaliseAngle(angle float64) float64 {
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

func AdjustTo360(degrees float64) float64 {
	for degrees < 0.0 {
		degrees = degrees + 360.0
	}
	for degrees > 360.0 {
		degrees = degrees - 360.0
	}
	return degrees
}
