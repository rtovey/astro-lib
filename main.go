package main

import (
	//"astro/solar"
	"fmt"
	"time"

	c "./common"
	"./lunar"
)

func main() {
	//orbitRoutineDate := time.Date(1988, time.July, 27, 0, 0, 0, 0, time.UTC)
	//solar.SunRiseAndSet(orbitRoutineDate)
	//sunRiseRoutineDate := time.Date(1986, time.March, 10, 0, 0, 0, 0, time.UTC)
	//solar.SunRiseAndSet(sunRiseRoutineDate)

	// date := time.Now()
	// end := date.Add(24 * time.Hour)

	// for end.Sub(date) > 0 {
	// 	date = date.Add(time.Minute)

	// 	phase := lunar.Phase(date)

	// 	fmt.Printf("%s = %.0f%%\n", date.Format("2006-01-02 15:04"), phase*100)
	// }

	phase := lunar.Phase(time.Now())
	fmt.Printf("Lunar phase: %.0f%%\n", phase*100)

	lunar.RiseTime(time.Date(1979, time.February, 26, 16, 0, 50, 0, time.UTC))
	//lunar.RiseTime(time.Date(1986, time.March, 6, 0, 0, 0, 0, time.UTC))

	ec := c.Ecliptic{
		Longitude: 139.686111,
		Latitude:  4.875278,
	}
	eq := ec.ToEquatorial(time.Now())
	fmt.Printf("%+v\n\n", eq)
}
