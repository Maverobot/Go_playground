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

func safeIncrease(v byte) (byte, bool) {
	if v < 255 {
		return v + 1, false
	}
	return 0, true
}

// RGB contains R, G, B values of a color representation
type RGB struct {
	r, g, b byte
}

// Next increases RGB value by 1, in the order of R <- G <- B
func (rgb *RGB) Next() {
	var zeroed bool
	rgb.b, zeroed = safeIncrease(rgb.b)
	if zeroed {
		rgb.g, zeroed = safeIncrease(rgb.g)
		if zeroed {
			rgb.r, _ = safeIncrease(rgb.r)
		}
	}
}

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	ledR := gpio.NewLedDriver(firmataAdaptor, "6")
	ledG := gpio.NewLedDriver(firmataAdaptor, "5")
	ledB := gpio.NewLedDriver(firmataAdaptor, "3")

	rgb := &RGB{
		r: 0,
		g: 0,
		b: 0,
	}

	work := func() {
		gobot.Every(time.Microsecond, func() {
			fmt.Printf("rgb = [%d, %d, %d]\n", rgb.r, rgb.g, rgb.b)
			ledR.Brightness(rgb.r)
			ledG.Brightness(rgb.g)
			ledB.Brightness(rgb.b)
			rgb.Next()
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
