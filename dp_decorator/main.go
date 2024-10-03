package main

import "fmt"

func main() {

	pizza := &VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndVegetables := &VegetableTopping{
		pizza: pizzaWithCheese,
	}

	pizzarouge := &BaseRouge{}

	//Add anchovy topping
	pizzaWithAnchovy := &AnchovyTopping{
		pizza: pizzarouge,
	}

	pizzablanche := &BaseBlanche{}

	pizzaWithHam := &HamTopping{
		pizza: pizzablanche,
	}

	pizzaWithHamAndMushroom := &MushroomTopping{
		pizza: pizzaWithHam,
	}

	fmt.Printf("Le prix de la %s est %d €\n", pizzaWithCheeseAndVegetables.getName(), pizzaWithCheeseAndVegetables.getPrice())
	fmt.Printf("Le prix de la %s est %d €\n", pizzaWithAnchovy.getName(), pizzaWithAnchovy.getPrice())
	fmt.Printf("Le prix de la %s est %d €\n", pizzaWithHamAndMushroom.getName(), pizzaWithHamAndMushroom.getPrice())
}
