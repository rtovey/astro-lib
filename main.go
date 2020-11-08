package main

import (
	//"astro/solar"

	"time"

	c "github.com/rtovey/astro/common"
	"github.com/rtovey/astro/solar"

	//"github.com/rtovey/astro/lunar"

	"github.com/davecgh/go-spew/spew"
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

	date := time.Now()
	// Greenwich Observatory, London, UK
	loc, _ := time.LoadLocation("Europe/London")
	observer := c.Observer{
		Latitude:  51.477840,
		Longitude: 0.0,
		Location:  loc,
	}

	/*
		// Worked example
		loc, _ := time.LoadLocation("EST")
		date := time.Date(1986, time.March, 6, 0, 0, 0, 0, loc)
		observer := c.Observer{
			Latitude:  42.3666667,
			Longitude: -71.05,
			Location:  loc,
		}*/

	/*lunarPhase := lunar.Phase(date)
	lunarRiseSetTime := lunar.RiseSetTime(observer, date)

	spew.Dump(lunarRiseSetTime)
	fmt.Printf("\n\n\nLunar phase: %.0f%%\nRise time: %v\nSet time: %v\n", lunarPhase*100, lunarRiseSetTime.Rise, lunarRiseSetTime.Set)*/

	solarRiseSetTime := solar.RiseSetTime(observer, date)

	spew.Dump(solarRiseSetTime)

}
