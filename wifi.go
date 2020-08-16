package gopherboat

import (
	"machine"

	"tinygo.org/x/drivers/wifinina"
)

// WiFiDevice controls the WiFi on Gopherboat.
type WiFiDevice struct {
	wifinina.Device
}

// NewWiFiDevice returns a new WiFi device.
func NewWiFiDevice() *WiFiDevice {
	EnsureSPIInit()

	adaptor := wifinina.Device{
		SPI:   machine.SPI0,
		CS:    WIFI_CS,
		ACK:   WIFI_ACK,
		GPIO0: WIFI_GPIO0,
		RESET: WIFI_RESETN,
	}

	adaptor.Configure()

	return &WiFiDevice{
		Device: adaptor,
	}
}
