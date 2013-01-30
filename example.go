//
// blink(1) USB interface demo
//
// Derived from https://github.com/GaryBoone/GoBlink by Gary Boone. 
// Modified by: Nick Granado, 2013
//
// To run type:
// $ go run example.go
//

package main

import "github.com/heatxsink/blink1-go/blink1"

import (
	"time"
)

func main() {
	blink_interval := 250 * time.Millisecond
	time_one_second := 1 * time.Second
	time_ten_seconds := 10 * time.Second
		
	b := blink1.NewBlink()
	
	// first lets blink 10 times
	b.Blink(10, blink_interval)
	
	time.Sleep(time_one_second)
	
	// second lets set the blink(1) to a specific color = Go Giants!
	b.SetRGB(255, 114, 0)
	
	time.Sleep(time_ten_seconds)
	
	// third just choose 20 random colors per second
	b.Random(20, time_one_second)
}
