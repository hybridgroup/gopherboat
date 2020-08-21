package gopherboat

import (
	"machine"
)

const (
	LeftDirPin   = machine.D12
	LeftBrakePin = machine.D9
	LeftSpeedPin = machine.D3

	RightDirPin   = machine.D13
	RightBrakePin = machine.D8
	RightSpeedPin = machine.D11
)

// ShieldMotorDevice controls the motors on Gopherboat.
type ShieldMotorDevice struct {
	dir, brake machine.Pin
	speed      machine.PWM
}

// NewShieldMotor returns a new motor device.
func NewShieldMotor(dir, brake, speed machine.Pin) *ShieldMotorDevice {
	dir.Configure(machine.PinConfig{Mode: machine.PinOutput})
	brake.Configure(machine.PinConfig{Mode: machine.PinOutput})
	s := machine.PWM{speed}
	s.Configure()

	return &ShieldMotorDevice{
		dir:   dir,
		brake: brake,
		speed: s,
	}
}

// NewLeftShieldMotor returns the left motor device.
func NewLeftShieldMotor() *ShieldMotorDevice {
	return NewShieldMotor(LeftDirPin, LeftBrakePin, LeftSpeedPin)
}

// NewRightShieldMotor returns the right motor device.
func NewRightShieldMotor() *ShieldMotorDevice {
	return NewShieldMotor(RightDirPin, RightBrakePin, RightSpeedPin)
}

// Forward goes forward.
func (m *ShieldMotorDevice) Forward(s uint16) {
	m.brake.Low()
	m.dir.High()
	m.speed.Set(s)
}

// Backward goes backward.
func (m *ShieldMotorDevice) Backward(s uint16) {
	m.brake.Low()
	m.dir.Low()
	m.speed.Set(s)
}

// Stop the motor.
func (m *ShieldMotorDevice) Stop() {
	m.speed.Set(0)
	m.brake.High()
}
