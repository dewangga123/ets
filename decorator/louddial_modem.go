package decorator

type LoudDialModem struct {
	*DecoratorModem
}

func (modem *LoudDialModem) Dial(pno string) {
	modem.DecoratorModem.SetSpeakerVolume(10)
	modem.DecoratorModem.Dial(pno)
}
