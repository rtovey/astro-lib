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

func Tand(degrees float64) float64 {
	return math.Tan(dtoR(degrees))
}

func Asind(x float64) float64 {
	return rtoD(math.Asin(x))
}

func Acosd(x float64) float64 {
	return rtoD(math.Acos(x))
}

func Atan2d(x float64, y float64) float64 {
	atan := rtoD(math.Atan2(x, y))
	return AdjustTo360(atan)
}

func DtoR(degrees float64) float64 {
	return dtoR(degrees)
}

func dtoR(degrees float64) float64 {
	return (math.Pi / 180) * degrees
}

func RtoD(radians float64) float64 {
	return rtoD(radians)
}

func rtoD(radians float64) float64 {
	return (180 * radians) / math.Pi
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
