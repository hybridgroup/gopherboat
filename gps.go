package gopherboat

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/gps"
)

// GPSDevice controls the GPS on Gopherboat.
type GPSDevice struct {
	gps.Device
	Fixes  chan *gps.Fix
	Errors chan error
}

// NewGPSDevice returns a new GPS device.
func NewGPSDevice() *GPSDevice {
	EnsureUARTInit()
	ublox := gps.NewUART(&machine.UART1)

	return &GPSDevice{
		Device: ublox,
		Fixes:  make(chan *gps.Fix, 10),
		Errors: make(chan error, 10),
	}
}

// Start getting GPS coordinates. Best run as a goroutine.
func (gpsu *GPSDevice) Start() {
	parser := gps.NewParser()
	for {
		s, err := gpsu.NextSentence()
		if err != nil {
			gpsu.Errors <- err
			continue
		}

		fix, err := parser.Parse(s)
		if err != nil {
			gpsu.Errors <- err
			continue
		}

		if fix.Valid {
			gpsu.Fixes <- &fix
		}

		time.Sleep(100 * time.Millisecond)
	}
}
