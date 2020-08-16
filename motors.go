package gopherboat

import (
	"machine"

	"tinygo.org/x/drivers/l293x"
)

// MotorDevice controls the motors on Gopherboat.
type MotorDevice struct {
	l293x.PWMDevice
}

// NewMotor returns a new motor device.
func NewMotor(a1, a2, en machine.Pin) *MotorDevice {
	mtr := l293x.NewWithSpeed(a1, a2, machine.PWM{en})

	return &MotorDevice{
		PWMDevice: mtr,
	}
}

// NewLeftMotor returns the left motor device.
func NewLeftMotor() *MotorDevice {
	return nil
}

// NewRightMotor returns the right motor device.
func NewRightMotor() *MotorDevice {
	return nil
}
