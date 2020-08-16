package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gopherboat"
)

func main() {
	gpsunit := gopherboat.NewGPSDevice()

	go gpsunit.Start()

	for {
		fix := gpsunit.Fix
		if fix.Valid {
			print(fix.Time.Format("15:04:05"))
			print(", lat=", fmt.Sprintf("%f", fix.Latitude))
			print(", long=", fmt.Sprintf("%f", fix.Longitude))
			print(", altitude:=", fix.Altitude)
			print(", satellites=", fix.Satellites)
			println()
		} else {
			println("No fix")
		}

		time.Sleep(1 * time.Second)
	}
}
