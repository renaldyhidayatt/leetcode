package bridge

import (
	"fmt"
	"structuralpatterns/bridge/models"
)

func Bridge() {
	hpPrinter := &models.Hp{}
	epsonPrinter := &models.Epson{}
	macComputer := &models.Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()
	winComputer := &models.Windows{}
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()
	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
