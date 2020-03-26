package decorator

type DecoratorModem struct {
	Modem Modem
}

func (modem *DecoratorModem) Dial(pno string) {
	modem.Modem.Dial(pno)
}
func (modem *DecoratorModem) SetSpeakerVolume(volume int) {
	modem.Modem.SetSpeakerVolume(volume)
}
func (modem *DecoratorModem) GetPhoneNumber() string {
	return modem.Modem.GetPhoneNumber()
}
func (modem *DecoratorModem) GetSpeakerVolume() int {
	return modem.Modem.GetSpeakerVolume()
}
