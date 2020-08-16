// Package gopherboat is the TinyGo code for the Gopherboat programmable robotic boat.
package gopherboat

import "machine"

var (
	pwmInitComplete = false
	i2cInitComplete = false
	spiInitComplete = false
)

// EnsurePWMInit makes sure that the Gopherboat PWM has been initialized, but
// is only initialized once.
func EnsurePWMInit() {
	if !pwmInitComplete {
		machine.InitPWM()
		pwmInitComplete = true
	}
}

// EnsureI2CInit makes sure that the Gopherboat I2C has been initialized, but
// is only initialized once.
func EnsureI2CInit() {
	if !i2cInitComplete {
		machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})
		i2cInitComplete = true
	}
}

// EnsureSPIInit makes sure that the Gopherboat SPI has been initialized, but
// is only initialized once.
func EnsureSPIInit() {
	if !spiInitComplete {
		// Configure SPI for 8Mhz, Mode 0, MSB First
		machine.SPI0.Configure(machine.SPIConfig{
			Frequency: 8 * 1e6,
			SDO:       machine.NINA_SDO,
			SDI:       machine.NINA_SDI,
			SCK:       machine.NINA_SCK,
		})
		spiInitComplete = true
	}
}
