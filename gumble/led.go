package gumble // import "github.com/talkkonnect/gumble/gumble"

import (
        "fmt"
        "github.com/stianeikeland/go-rpio"
        "os"
)

var (
        // Use mcu pin 10, corresponds to physical pin 19 on the pi
        //pin = rpio.Pin(10)
        pin = rpio.Pin(4)
)

func led(state string) {
        // Open and map memory to access gpio, check for errors
        if err := rpio.Open(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        // Unmap gpio memory when done
        defer rpio.Close()

        // Set pin to output mode
        pin.Output()

	if state =="on" {
	        pin.High()
	}
	if state =="off" {
	        pin.Low()
	}

}


