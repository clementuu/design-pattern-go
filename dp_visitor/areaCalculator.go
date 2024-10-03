package main

import (
	"fmt"
	"math"
)

type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
	a.area = float64(s.side) * float64(s.side)
	fmt.Println(a.area)
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
	a.area = float64(s.radius) * math.Pi * math.Pi
	fmt.Println(a.area)
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
	a.area = float64(s.b) * float64(s.l)
	fmt.Println(a.area)
}
