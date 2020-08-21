# Gopherboat

TinyGo-powered bottle boat.

## Vessel

See the "Hull" page at [assembly/hull.md](./assembly/hull.md) for more information about putting together the hull of your Gopherboat.

## Controller

- Microcontroller, such as Adafruit Metro M4 Airlift (ATSAMD51 microcontroller with onboard ESP32 WiFi)
- WiFi co-processor such as ESP32 (already onboard the Adafruit Metro M4 Airlift)
- Arduino Motor Controller
- 2 DC Motors
- 2 propellers (clockwise, 1 counterclockwise)
- Adafruit Ultimate GPS UART
- Adafruit I2C Digital Compass

See the "Wiring" page at [assembly/wiring.md](./assembly/wiring.md) for more information about wiring up your Gopherboat.

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

## Getting started

- Install Go 1.15

You probably already have this. If not, go install Go 1.15 from here: https://golang.org/dl/

- Install TinyGo 0.14.1

On MacOS you can install TinyGo using Homebrew like this:

```
brew tap tinygo-org/tools
brew install tinygo
```

For more info, see the "quick install" instructions here: https://tinygo.org/getting-started/macos/#quick-install

## Testing that you can flash the board

You use the `tinygo flash` command to put new software on the Adafruit Metro M4 Airlift board. First plug in your computer's USB port to the Adafruit Metro M4 Airlift board's micro-USB connector.

```
tinygo flash -target metro-m4-airlift examples/blinky1
```

This should flash the "blinky" program onto your Metro M4 board. The built-in LED should start blinking every 1/2 second.

Now you are ready!

## Examples

Take a look at the "Examples" directory for examples of how to use the various hardware.
