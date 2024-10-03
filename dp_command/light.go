package main

import "fmt"

type Light struct {
	isTurnOn bool
}

func (l *Light) on() {
	l.isTurnOn = true
	fmt.Println("Light on !")
}

func (l *Light) off() {
	l.isTurnOn = false
	fmt.Println("Light off !")
}
