package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	present := time.Now() //current Time
	p(present)

	DOB := time.Date(1993, 02, 28, 9, 01, 39, 213, time.Local)
	p(DOB)

	p(DOB.Year())
	p(DOB.Month())
	p(DOB.Day())

	p(DOB.Hour())
	p(DOB.Minute())
	p(DOB.Second())

	p(DOB.Nanosecond())
	p(DOB.Location())

	p(DOB.Weekday())

	p(DOB.Before(present))
	p(DOB.After(present))
	p(DOB.Equal(present))

	diff := present.Sub(DOB)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(DOB.Add(diff))
	p(DOB.Add(-diff))

}
