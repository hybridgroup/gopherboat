# Wiring

## A note about color coding

It is best to try to always use the same "standard" colored jumper cables to avoid confusion when plugging things in. For example, try to always only use black cables for the "Ground" (G). Also, only use red cables for power.

## Main board

Start with the Adafruit Metro M4 Airlift board.

Connect the Arduino Motor Shield by plugging in into the top of the Adafruit Metro M4 Airlift. Be careful not to bend the pins.

## GPS

Plug in the GPS board to the Arduino Motor Shield using the male to female jumper cables as follows:

- Plug the GPS "G" to the Arduino Motor Shield "G"
- Plug the GPS "VIN" to the Arduino Motor Shield "5V"
- Plug the GPS "TX" to the Arduino Motor Shield "RX" (Pin 0)
- Plug the GPS "RX" to the Arduino Motor Shield "TX" (Pin 1)

## Digital compass

Plug in the LSM303AGR digital compass board to the Arduino Motor Shield using the male to female jumper cables as follows:

- Plug the LSM303AGR "G" to the Arduino Motor Shield "G"
- Plug the LSM303AGR "3V" to the Arduino Motor Shield "3V"
- Plug the LSM303AGR "SCL" to the Arduino Motor Shield "SCL" (pin nearest end of board)
- Plug the LSM303AGR "SDA" to the Arduino Motor Shield "SDA" (pin right next to "SCL" pin)

## Motors

Connect the wires from the DC motors to the screw terminal blocks on the Arduino Motor Shield.

We will need more details here.

## Power

Plug in the barrel connector from the battery pack into the plug on the Metro M4 board. This should provide power for all of the components of the system, including the motors.
