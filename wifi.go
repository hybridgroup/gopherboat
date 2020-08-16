package gopherboat

import (
	"machine"
	"time"

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

// ConnectToAP will connect to access point.
func (wifi *WiFiDevice) ConnectToAP(ssid, pass string) {
	// wait for wifi hardware to startup if connecting
	time.Sleep(2 * time.Second)

	println("Connecting to " + ssid)
	wifi.SetPassphrase(ssid, pass)
	for st, _ := wifi.GetConnectionStatus(); st != wifinina.StatusConnected; {
		println("Connection status: " + st.String())
		time.Sleep(1 * time.Second)
		st, _ = wifi.GetConnectionStatus()
	}
	println("Connected.")
	time.Sleep(2 * time.Second)
	ip, _, _, err := wifi.GetIP()
	for ; err != nil; ip, _, _, err = wifi.GetIP() {
		println(err.Error())
		time.Sleep(1 * time.Second)
	}
	println(ip.String())
}
