package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hybridgroup/gopherboat"
	"tinygo.org/x/drivers/net/mqtt"
	"tinygo.org/x/drivers/wifinina"
)

const (
	ssid   = "YOURSSID"
	pass   = "YOURPASS"
	server = "192.168.0.1"
	topic  = "gopherboat"
)

func main() {
	wifi := gopherboat.NewWiFiDevice()
	wifi.ConnectToAP(ssid, pass)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(server).SetClientID("gopherboat-" + randomString(10))

	println("Connecting to MQTT broker at", server)
	cl := mqtt.NewClient(opts)
	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		failMessage(token.Error().Error())
	}

	i := 0

	for {
		println("Publishing MQTT message...")
		data := []byte(fmt.Sprintf(`{"e":[{"n":"hello %d","v":101}]}`, i))
		token := cl.Publish(topic, 0, false, data)
		token.Wait()
		if err := token.Error(); err != nil {
			switch t := err.(type) {
			case wifinina.Error:
				println(t.Error(), "attempting to reconnect to mqtt")
				if token := cl.Connect(); token.Wait() && token.Error() != nil {
					failMessage(token.Error().Error())
				}
			default:
				println(err.Error())
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

// output fail message via USBCDC
func failMessage(msg string) {
	for {
		println(msg)
		time.Sleep(1 * time.Second)
	}
}
