package adapter

import "structuralpatterns/adapter/models"

func Adapter() {
	client := &models.Client{}

	mac := &models.Mac{}

	client.InsertSquareUsbInComputer(mac)
	windowsMachine := &models.Windows{}
	windowsMachineAdapter := &models.WindowsAdapter{
		WindowMachine: windowsMachine,
	}
	client.InsertSquareUsbInComputer(windowsMachineAdapter)
}
