package cyclic_visitor

type HayesModem struct {
	ConfigurationString string
}

func (modem *HayesModem) Dial(pno string) {}
func (modem *HayesModem) Hangup()         {}
func (modem *HayesModem) Send(c rune)     {}
func (modem *HayesModem) Recv()           {}
func (modem *HayesModem) Accept(modemVisitor ModemVisitor) {
	if visitor, ok := modemVisitor.(HayesVisitor); ok {
		visitor.Visit(modem)
	}
}

func (modem *HayesModem) GetConfiguration() interface{} {
	return modem.ConfigurationString
}
