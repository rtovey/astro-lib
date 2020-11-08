package common

import "math"

func SolveKeplersEquation(M float64, e float64, accuracy float64) float64 {
	E := M
	d := accuracy * 10
	for d > accuracy {
		d = E - (e * math.Sin(E)) - M
		dE := d / (1 - (e * math.Cos(E)))
		E = E - dE
	}
	return E
}
