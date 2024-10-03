package main

import "fmt"

type Radio struct {
	isRunning bool
}

func (r *Radio) on() {
	r.isRunning = true
	fmt.Println("Turning the radio on")
}

func (r *Radio) off() {
	r.isRunning = false
	fmt.Println("Turning the radio off")
}
