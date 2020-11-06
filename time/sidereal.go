package time

import (
	c "../common"
)

type LocalSiderealTime struct {
	Time float64
}

func (lst LocalSiderealTime) ToGst(observer c.Observer) float64 {
	gst := lst.Time
	timeDiff := observer.Longitude / 15.0
	if observer.Longitude < 0 {
		gst += timeDiff
	}
	if observer.Longitude > 0 {
		gst -= timeDiff
	}
	if gst > 24.0 {
		gst -= 24.0
	}
	if gst < 0 {
		gst += 24.0
	}
	return gst
}
