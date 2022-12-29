package models

import (
	"fmt"
	"structuralpatterns/bridge/interfaces"
)

type Windows struct {
	Printer interfaces.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.Printer.PrintFile()
}

func (w *Windows) SetPrinter(p interfaces.Printer) {
	w.Printer = p
}
