package orbit

import c "../common"

func AngleAtHorizon(observer c.Observer, position Equatorial) float64 {
	return c.Acosd(c.Sind(observer.Latitude) / c.Cosd(position.Declination))
}

func Y(x float64, angleAtHorizon float64) float64 {
	return c.Asind(c.Sind(x) / c.Sind(angleAtHorizon))
}

func RiseSetTimeShiftSeconds(y float64, position Equatorial) float64 {
	return (240.0 * y) / c.Cosd(position.Declination)
}
