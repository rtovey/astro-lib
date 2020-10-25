package lunar

import (
	"fmt"
	"time"
)

func RiseTime(date time.Time) time.Time {
	position := Position(date)
	fmt.Printf("%+v\n", position)
	return time.Now()
}
