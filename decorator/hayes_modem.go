package decorator

type HayesModem struct {
	PhoneNumber string
	Volume      int
}

func (modem *HayesModem) Dial(pno string) {
	modem.PhoneNumber = pno
}
func (modem *HayesModem) SetSpeakerVolume(volume int) {
	modem.Volume = volume
}
func (modem *HayesModem) GetPhoneNumber() string {
	return modem.PhoneNumber
}
func (modem *HayesModem) GetSpeakerVolume() int {
	return modem.Volume
}
