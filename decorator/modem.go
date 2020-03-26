package decorator

type Modem interface {
	Dial(pno string)
	SetSpeakerVolume(volume int)
	GetPhoneNumber() string
	GetSpeakerVolume() int
}
