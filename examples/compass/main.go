package main

import (
	"time"

	"github.com/hybridgroup/gopherboat"
)

func main() {
	compass := gopherboat.NewCompassDevice()

	for {
		heading := compass.ReadCompass()
		println(heading)

		time.Sleep(1 * time.Second)
	}
}
