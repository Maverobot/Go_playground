/*
 How to run
 Pass serial port to use as the first param:

	go run arduino_digital_input.go /dev/ttyACM0
*/

package main

import (
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])

	button := gpio.NewButtonDriver(firmataAdaptor, "5")
	led := gpio.NewLedDriver(firmataAdaptor, "3")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{button, led},
		work,
	)

	robot.Start()
}
