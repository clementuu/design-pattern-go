package main

import "fmt"

type AirConditionner struct {
	isRunning bool
}

func (a *AirConditionner) on() {
	a.isRunning = true
	fmt.Println("Turning the air conditionner on")
}

func (a *AirConditionner) off() {
	a.isRunning = false
	fmt.Println("Turning the air conditionner off")
}
