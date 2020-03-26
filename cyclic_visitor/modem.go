package cyclic_visitor

type Modem interface {
	Dial(pno string)
	Hangup()
	Send(c rune)
	Recv()
	Accept(modemVisitor ModemVisitor)
}
