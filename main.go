package main

import (
	//"astro/solar"
	"fmt"
	"time"

	"./lunar"
)

func main() {
	//orbitRoutineDate := time.Date(1988, time.July, 27, 0, 0, 0, 0, time.UTC)
	//solar.SunRiseAndSet(orbitRoutineDate)
	//sunRiseRoutineDate := time.Date(1986, time.March, 10, 0, 0, 0, 0, time.UTC)
	//solar.SunRiseAndSet(sunRiseRoutineDate)

	phase := lunar.Phase(time.Now())
	fmt.Printf("Lunar phase: %.0f%%\n", phase*100)

	lunar.RiseTime(time.Date(1979, time.February, 26, 16, 0, 50, 0, time.UTC))
}
