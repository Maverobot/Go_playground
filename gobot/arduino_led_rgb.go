/*
 How to run
 Pass serial port to use as the first param:

	go run arduino_led_rgb.go /dev/ttyACM0
*/

package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	ledR := gpio.NewLedDriver(firmataAdaptor, "6")
	ledG := gpio.NewLedDriver(firmataAdaptor, "5")
	ledB := gpio.NewLedDriver(firmataAdaptor, "3")

	leds := []*gpio.LedDriver{ledR, ledG, ledB}

	var count int
	count = 0
	work := func() {
		gobot.Every(2*time.Second, func() {
			for i, led := range leds {
				if i != count {
					led.Off()
				} else {
					led.On()
				}
			}
			count++
			count = count % len(leds)
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		//		[]gobot.Device{leds[0]},
		//		[]gobot.Device{leds[1]},
		//		[]gobot.Device{leds[2]},
		work,
	)

	err := robot.Start()
	if err != nil {
		fmt.Println(err)
	}
}
