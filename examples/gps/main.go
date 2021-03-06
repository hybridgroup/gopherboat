package main

import (
	"time"

	"github.com/hybridgroup/gopherboat"
)

func main() {
	gpsunit := gopherboat.NewGPSDevice()

	go gpsunit.Start()

	for {
		select {
		case fix := <-gpsunit.Fixes:
			if fix.Valid {
				print(fix.Time.Format("15:04:05"))
				print(", lat=")
				print(fix.Latitude)
				print(", long=")
				print(fix.Longitude)
				print(", altitude=", fix.Altitude)
				print(", satellites=", fix.Satellites)
				if fix.Speed != 0 {
					print(", speed=")
					print(fix.Speed)
				}
				if fix.Heading != 0 {
					print(", heading=")
					print(fix.Heading)
				}
				println()
			} else {
				println("No fix")
			}
		case err := <-gpsunit.Errors:
			println(err)

		default:
			println("No fix")
		}

		time.Sleep(100 * time.Millisecond)
	}
}
