/*
 How to run
 Pass serial port to use as the first param:

	go run arduino_direct_pin.go /dev/ttyACM0
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

	tiltBallSwitch := gpio.NewDirectPinDriver(firmataAdaptor, "12")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		var val int
		var err error
		for ; true; val, err = tiltBallSwitch.DigitalRead() {
			if val == 0 {
				// Turn on LED when the ball connects the pins
				err = led.On()
			} else if val == 1 {
				// Turn off LED when the ball is detached from the pins
				err = led.Off()
			}
			if err != nil {
				panic(err)
			}
		}
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{tiltBallSwitch, led},
		work,
	)

	err := robot.Start()
	if err != nil {
		panic(err)
	}
}
