package time

import (
	"math"
)

type Hours float64

func (h Hours) Value() float64 {
	return float64(h)
}

func (h Hours) Hours() int {
	f, _ := math.Modf(h.Value())
	return int(f)
}

func (h Hours) Minutes() int {
	_, f := math.Modf(h.Value())
	m, _ := math.Modf(f * 60.0)
	return int(m)
}

func (h Hours) Seconds() int {
	_, f := math.Modf(h.Value())
	_, g := math.Modf(f * 60.0)
	secs, _ := math.Modf(g * 60.0)
	return int(secs)
}
