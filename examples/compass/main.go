package main

import (
	"time"

	"github.com/hybridgroup/gopherboat"
)

func main() {
	compass := gopherboat.NewCompassDevice()
	time.Sleep(100 * time.Millisecond)

	for {
		heading := compass.ReadCompass()
		println("Heading: ", heading)

		time.Sleep(100 * time.Millisecond)
	}
}
