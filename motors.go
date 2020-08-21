package gopherboat

import (
	"machine"

	"tinygo.org/x/drivers/l293x"
)

const (
	LeftA1Pin = machine.D12
	LeftA2Pin = machine.D9
	LeftEnPin = machine.D3

	RightA1Pin = machine.D13
	RightA2Pin = machine.D8
	RightEnPin = machine.D11
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
	return NewMotor(LeftA1Pin, LeftA2Pin, LeftEnPin)
}

// NewRightMotor returns the right motor device.
func NewRightMotor() *MotorDevice {
	return NewMotor(RightA1Pin, RightA2Pin, RightEnPin)
}
