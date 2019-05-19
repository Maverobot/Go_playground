/*
 How to run
 Pass serial port to use as the first param:

	go run arduino_direct_pin.go /dev/ttyACM0
*/

package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	trigPin := gpio.NewDirectPinDriver(r, "12")
	echoPin := gpio.NewDirectPinDriver(r, "11")
	led := gpio.NewLedDriver(r, "13")

	work := func() {

		gobot.Every(1*time.Second, func() {

			println("Starting probing ")
			led.On()
			trigPin.DigitalWrite(byte(0))
			time.Sleep(2 * time.Microsecond)

			trigPin.DigitalWrite(byte(1))
			time.Sleep(10 * time.Microsecond)

			trigPin.DigitalWrite(byte(0))
			start := time.Now()
			end := time.Now()

			for {
				val, err := echoPin.DigitalRead()
				start = time.Now()

				if err != nil {
					println(err)
					break
				}

				if val == 0 {
					continue
				}

				break
			}

			for {
				val, err := echoPin.DigitalRead()
				end = time.Now()
				if err != nil {
					println(err)
					break
				}

				if val == 1 {
					continue
				}

				break
			}

			duration := end.Sub(start)
			durationAsInt64 := int64(duration)
			distance := duration.Seconds() * 34300
			distance = distance / 2 //one way travel time
			fmt.Printf("Duration : %v %v %v \n", distance, duration.Seconds(), durationAsInt64)
		})
	}

	robot := gobot.NewRobot("UltrasonicBot",
		[]gobot.Connection{r},
		[]gobot.Device{trigPin, echoPin, led},
		work,
	)

	robot.Start()
}
