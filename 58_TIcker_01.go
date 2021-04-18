package main

import (
	"fmt"
	"time"
)

func main() {
	tickerValue := time.NewTicker(time.Millisecond * 100)
	go func() {
		for t := range tickerValue.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 500)
	tickerValue.Stop()
	fmt.Println("Ticker Stopped.")
}
