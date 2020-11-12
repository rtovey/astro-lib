package solar

import (
	"math"

	c "github.com/rtovey/astro/common"
)

const (
	orbitalEccentricity   = 0.016713
	angularSizeAtApogee   = 0.533128    // degrees
	horizontalParallax    = 0.002441667 // degrees
	atmosphericRefraction = 0.56666667  // degrees
)

func apparentAngularDiameter(position SolarPosition) float64 {
	return angularSizeAtApogee * ((1 + (orbitalEccentricity * c.Cosd(position.Debug.v))) / (1 - math.Pow(orbitalEccentricity, 2)))
}
