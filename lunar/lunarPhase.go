package lunar

import (
	"time"

	c "../common"
)

func Phase(date time.Time) float64 {
	position := Position(date)

	F := lunarPhase(position.Debug.lll, position.Debug.Ls)
	return F
}

func lunarPhase(lll float64, L float64) float64 {
	return 0.5 * (1 - c.Cosd(lll-L))
}
