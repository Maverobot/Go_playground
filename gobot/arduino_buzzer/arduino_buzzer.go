/*
 How to run
 Pass serial port to use as the first param:

	go run arduino_buzzer.go /dev/ttyACM0
*/

package main

import (
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])

	buzzer := gpio.NewBuzzerDriver(firmataAdaptor, "12")

	work := func() {
		type note struct {
			tone     float64
			duration float64
		}

		song := []note{
			{gpio.C4, gpio.Quarter},
			{gpio.G3, gpio.Eighth},
			{gpio.G3, gpio.Eighth},
			{gpio.A3, gpio.Quarter},
			{gpio.G3, gpio.Quarter},
			{gpio.Rest, gpio.Quarter},
			{gpio.B3, gpio.Quarter},
			{gpio.C4, gpio.Quarter},
		}

		for _, val := range song {
			err := buzzer.Tone(val.tone, val.duration)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Duration(0.1 * val.duration * float64(time.Second)))
		}
	}

	robot := gobot.NewRobot("buzzerBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{buzzer},
		work,
	)

	err := robot.Start()
	if err != nil {
		panic(err)
	}
}
