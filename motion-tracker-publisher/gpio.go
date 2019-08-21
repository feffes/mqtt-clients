package main

import (
	"fmt"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

func poll(sensor embd.DigitalPin, ch chan string) {
	state := 0
	var err error
	for state != 1 {
		state, err = sensor.Read()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
	}
	fmt.Printf("Triggered")
	ch <- "Triggered"
}

func pinConnect(pin string) embd.DigitalPin {
	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()

	sensor, err := embd.NewDigitalPin(pin)
	if err != nil {
		panic(err)
	}
	defer sensor.Close()

	if err := sensor.SetDirection(embd.In); err != nil {
		panic(err)
	}
	return sensor
}
