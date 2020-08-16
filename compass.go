package gopherboat

import (
	"machine"

	"tinygo.org/x/drivers/lsm303agr"
)

// CompassDevice controls the compass on Gopherboat.
type CompassDevice struct {
	lsm303agr.Device
}

// NewCompassDevice returns a new compass device.
func NewCompassDevice() *CompassDevice {
	EnsureI2CInit()

	compass := lsm303agr.New(machine.I2C0)
	compass.Configure(lsm303agr.Configuration{})

	return &CompassDevice{
		Device: compass,
	}
}
