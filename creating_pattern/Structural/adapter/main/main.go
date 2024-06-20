package main

import "go-design-patterns/creating_pattern/Structural/adapter/lib"

func main() {

	client := &lib.Client{}
	mac := &lib.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &lib.Windows{}
	windowsMachineAdapter := &lib.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
