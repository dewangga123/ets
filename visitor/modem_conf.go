package visitor

type UnixModemConfigurator struct {
}

func (unix *UnixModemConfigurator) Visit(modem interface{}) {
	if input, ok := modem.(HayesModem); ok {
		input.ConfigurationString = "0"
	} else if input, ok := modem.(ErnieModem); ok {
		input.ConfigurationInt = 0
	} else if input, ok := modem.(ZoomModem); ok {
		input.ConfigurationFloat = 0.0
	}
}
