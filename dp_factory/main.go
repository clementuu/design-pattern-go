package main

import (
	"fmt"
)

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	gatling, _ := getGun("gatling")

	printDetails(ak47)
	printDetails(musket)
	printDetails(gatling)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
