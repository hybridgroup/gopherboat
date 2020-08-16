package gopherboat

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/gps"
)

// GPSDevice controls the GPS on Gopherboat.
type GPSDevice struct {
	gps.GPSDevice
	Fix gps.Fix
}

// NewGPSDevice returns a new GPS device.
func NewGPSDevice() *GPSDevice {
	ublox := gps.NewUART(&machine.UART1)

	return &GPSDevice{
		GPSDevice: ublox,
	}
}

// Start getting GPS coordinates. Best run as a goroutine.
func (gpsu *GPSDevice) Start() {
	parser := gps.Parser(gpsu.GPSDevice)
	var fix gps.Fix
	for {
		fix = parser.NextFix()
		if fix.Valid {
			gpsu.Fix = fix
		}
		time.Sleep(100 * time.Millisecond)
	}
}
