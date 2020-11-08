package lunar

import (
	"math"

	c "github.com/rtovey/astro/common"
)

const (
	orbitalEccentricity   = 0.0549
	angularSizeAtApogee   = 0.5181     // degrees
	parallaxAtApogee      = 0.9507     // degrees
	atmosphericRefraction = 0.56666667 // degrees
)

func apparentAngularDiameter(position LunarPosition) float64 {
	p := lunarSeparationFraction(position)
	return angularSizeAtApogee / p
}

func horizontalParallax(position LunarPosition) float64 {
	p := lunarSeparationFraction(position)
	return parallaxAtApogee / p
}

func lunarSeparationFraction(position LunarPosition) float64 {
	return (1 - math.Pow(orbitalEccentricity, 2)) / (1 + (orbitalEccentricity * c.Cosd(position.Debug.MMm+position.Debug.Ec)))
}
