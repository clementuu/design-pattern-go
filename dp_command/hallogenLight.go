package main

import "fmt"

type HallogenLight struct {
	isTurnOn bool
}

func (l *HallogenLight) on() {
	l.isTurnOn = true
	fmt.Println("Hallogen on !")
}

func (l *HallogenLight) off() {
	l.isTurnOn = false
	fmt.Println("Hallogen off !")
}
