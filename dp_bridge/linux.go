package main

import "fmt"

type Linux struct {
	printer Printer
}

func (l *Linux) Print() {
	fmt.Println("Print request for linux")
	l.printer.PrintFile()
}

func (l *Linux) SetPrinter(p Printer) {
	l.printer = p
}
