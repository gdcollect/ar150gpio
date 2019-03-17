package ar150gpio

import (
	"fmt"
	"os"
)

const IN uint8 = 0
const OUT uint8 = 1

const LOW uint8 = 0
const HIGH uint8 = 1

type gpio struct {
	pinNumber uint8
	pinMode uint8
	pinState uint8
}

func NewGPIO() gpio {
	return gpio {}
}

func (io *gpio) Initialize(pinNumber uint8, pinMode uint8) {
	(*io).pinNumber = pinNumber
	(*io).pinMode = pinMode
	(*io).export((*io).pinNumber)
	(*io).setMode((*io).pinMode)
}

func (io *gpio) export(pinNumber uint8) {
	f, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, 0755)
	check(err)

	defer f.Close()

	fmt.Fprintf(f, "%d", (*io).pinNumber)
}

func (io *gpio) setMode(pinMode uint8) {
	var directionFile string
	directionFile = fmt.Sprintf("/sys/class/gpio/gpio%d/direction", (*io).pinNumber)
	
	f, err := os.OpenFile(directionFile, os.O_WRONLY, 0755)
	check(err)

	defer f.Close()

	(*io).pinMode = pinMode

	if pinMode == IN {
		fmt.Fprintf(f, "in")
	}else {
		fmt.Fprintf(f, "out")
	}		
}

func (io *gpio) Write(pinState uint8) {
	var valueFile string
	valueFile = fmt.Sprintf("/sys/class/gpio/gpio%d/value", (*io).pinNumber)

	f, err := os.OpenFile(valueFile, os.O_WRONLY, 0755)
	check(err)

	defer f.Close()

	(*io).pinState = pinState;

	if pinState == LOW {
		fmt.Fprintf(f, "0")
	}else {
		fmt.Fprintf(f, "1")
	}

}

func (io *gpio) Read() uint8 {
	var valueFile string
	valueFile = fmt.Sprintf("/sys/class/gpio/gpio%d/value", (*io).pinNumber)

	f, err := os.OpenFile(valueFile, os.O_RDONLY, 0755)
	check(err)

	defer f.Close()

	fmt.Fscanf(f, "%d", &((*io).pinState))

	return (*io).pinState

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
