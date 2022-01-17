package lunar

import (
	"time"

	c "github.com/rtovey/astro/common"
)

type LunarPhase struct {
	Illuminated_pc float64
}

func Phase(date time.Time) LunarPhase {
	position := Position(date)

	F := lunarPhase(position.Debug.lll, position.Debug.Ls)
	return LunarPhase{
		Illuminated_pc: F,
	}
}

func lunarPhase(lll float64, L float64) float64 {
	return 0.5 * (1 - c.Cosd(lll-L))
}
