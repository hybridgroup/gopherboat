package gopherboat

import (
	"machine"

	"tinygo.org/x/drivers/lis2mdl"
)

// CompassDevice controls the compass on Gopherboat.
type CompassDevice struct {
	lis2mdl.Device
}

// NewCompassDevice returns a new compass device.
func NewCompassDevice() *CompassDevice {
	EnsureI2CInit()

	compass := lis2mdl.New(machine.I2C0)
	compass.Configure(lis2mdl.Configuration{})

	return &CompassDevice{
		Device: compass,
	}
}
