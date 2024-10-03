package main

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	a.x = s.side
	a.y = 0
	fmt.Println("Calculating middle point coordinates for square")
	fmt.Println(a.x, a.y)
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	a.x = c.radius
	a.y = 0
	fmt.Println("Calculating middle point coordinates for circle")
	fmt.Println(a.x, a.y)
}
func (a *MiddleCoordinates) visitForrectangle(t *Rectangle) {
	a.x = t.b
	a.y = t.l
	fmt.Println("Calculating middle point coordinates for rectangle")
	fmt.Println(a.x, a.y)
}
