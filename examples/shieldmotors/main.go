package main

import (
	"time"

	"github.com/hybridgroup/gopherboat"
)

const maxSpeed = 1000

func main() {
	leftmotor := gopherboat.NewLeftShieldMotor()
	rightmotor := gopherboat.NewRightShieldMotor()

	for i := 0; i < 10; i++ {
		leftmotor.Forward(maxSpeed)
		rightmotor.Backward(maxSpeed)

		time.Sleep(5 * time.Second)

		leftmotor.Backward(maxSpeed)
		rightmotor.Forward(maxSpeed)

		time.Sleep(5 * time.Second)
	}

	leftmotor.Stop()
	rightmotor.Stop()
}
