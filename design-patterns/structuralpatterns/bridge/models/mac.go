package models

import (
	"fmt"
	"structuralpatterns/bridge/interfaces"
)

type Mac struct {
	Printer interfaces.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(p interfaces.Printer) {
	m.Printer = p
}
