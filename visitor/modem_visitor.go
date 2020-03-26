package visitor

type ModemVisitor interface {
	Visit(modem interface{})
}
