package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hybridgroup/gopherboat"
	"github.com/hybridgroup/gopherboat/connect"
	"tinygo.org/x/drivers/net/mqtt"
)

const (
	ssid    = connect.SSID
	pass    = connect.PASS
	server  = connect.MQTTBroker
	topicTx = "gopherboat/tx"
	topicRx = "gopherboat/rx"
)

var cl mqtt.Client

func subHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[%s]  ", msg.Topic())
	fmt.Printf("%s\r\n", msg.Payload())
}

func main() {
	wifi := gopherboat.NewWiFiDevice()
	wifi.ConnectToAP(ssid, pass)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(server).SetClientID("gopherboat-" + randomString(10))

	println("Connecting to MQTT broker at", server)
	cl = mqtt.NewClient(opts)
	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		failMessage(token.Error().Error())
	}

	// subscribe
	token := cl.Subscribe(topicRx, 0, subHandler)
	token.Wait()
	if token.Error() != nil {
		failMessage(token.Error().Error())
	}

	go publishing()

	select {}
}

func publishing() {
	for i := 0; ; i++ {
		println("Publishing MQTT message...")
		data := []byte(fmt.Sprintf(`{"e":[{"n":"hello %d","v":101}]}`, i))
		token := cl.Publish(topicTx, 0, false, data)
		token.Wait()
		if token.Error() != nil {
			println(token.Error().Error())
		}

		time.Sleep(1000 * time.Millisecond)
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
