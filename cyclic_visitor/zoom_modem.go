package cyclic_visitor

type ZoomModem struct {
	ConfigurationFloat float32
}

func (modem *ZoomModem) Dial(pno string) {}
func (modem *ZoomModem) Hangup()         {}
func (modem *ZoomModem) Send(c rune)     {}
func (modem *ZoomModem) Recv()           {}
func (modem *ZoomModem) Accept(modemVisitor ModemVisitor) {
	if visitor, ok := modemVisitor.(ZoomVisitor); ok {
		visitor.Visit(modem)
	}
}

func (modem *ZoomModem) GetConfiguration() interface{} {
	return modem.ConfigurationFloat
}
