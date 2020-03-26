package main

import (
	"ets/decorator"
	"fmt"
)

func main() {
	hayesModem := &decorator.HayesModem{
		PhoneNumber: "1019191",
	}
	decoratorModem := &decorator.DecoratorModem{
		Modem: hayesModem,
	}
	loudDialModem := &decorator.LoudDialModem{
		DecoratorModem: decoratorModem,
	}
	GetDecorator(loudDialModem)

}

func GetDecorator(modem decorator.Modem) {
	modem.Dial("1019191")
	fmt.Println(modem.GetSpeakerVolume())
}
