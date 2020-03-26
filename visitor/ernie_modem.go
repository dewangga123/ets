package visitor

type ErnieModem struct {
	ConfigurationInt int
}

func (modem *ErnieModem) Dial(pno string) {}
func (modem *ErnieModem) Hangup()         {}
func (modem *ErnieModem) Send(c rune)     {}
func (modem *ErnieModem) Recv()           {}
func (modem *ErnieModem) Accept(modemVisitor ModemVisitor) {
	modemVisitor.Visit(modem)
}

func (modem *ErnieModem) GetConfiguration() interface{} {
	return modem.ConfigurationInt
}
