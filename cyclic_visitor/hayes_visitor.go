package cyclic_visitor

type HayesVisitor interface {
	Visit(modem *HayesModem)
}
