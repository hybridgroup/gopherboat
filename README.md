# Gopherboat

TinyGo-powered bottle boat.

## Vessel

- 2 or more empty plastic bottles for hull
- flat plastic sheet to use as deck
- risers for motors (must be able to lift motors/props high above water)
- Tie wraps
- Heavy duty transparent plastic bags
- Duct tape

## Controller

Microcontroller, such as Adafruit Metro M4 Airlift (ATSAMD51 microcontroller with onboard ESP32 WiFi)
WiFi co-processor such as ESP32 (already onboard the Adafruit Metro M4 Airlift)
Arduino Motor Controller
2 DC Motors
2 propellers (clockwise, 1 counterclockwise)
Adafruit Ultimate GPS UART
Adafruit I2C Digital Compass

## Code example

This code should make the Gopherboat pivot in place, first turning in a clockwise direction, and then turning back in a counter-clockwise direction.

```
package main

import (
	"time"

	"github.com/hybridgroup/gopherboat"
)

func main() {
	leftmotor := gopherboat.NewLeftMotor()
	rightmotor := gopherboat.NewRightMotor()

	for i := 0; i < 10; i++ {
		leftmotor.Forward()
		rightmotor.Backward()

		time.Sleep(5 * time.Second)

		leftmotor.Backward()
		rightmotor.Forward()

		time.Sleep(5 * time.Second)
	}

	leftmotor.Stop()
	rightmotor.Stop()
}
```
