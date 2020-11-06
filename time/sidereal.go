package time

import (
	c "../common"
)

type LST float64
type GST float64

func (lst LST) ToGst(observer c.Observer) float64 {
	gst := float64(lst)
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
