package lunar

import (
	"fmt"
	"time"
)

func RiseTime(date time.Time) time.Time {
	position := Position(date)
	fmt.Printf("%+v\n\n", position)
	equatorial := position.Ecliptic.ToEquatorial(date)
	fmt.Printf("%+v\n\n", equatorial)
	return time.Now()
}
