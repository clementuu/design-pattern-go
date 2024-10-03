package main

type AnchovyTopping struct {
	pizza IPizza
}

func (a *AnchovyTopping) getPrice() int {
	return a.pizza.getPrice() + 4
}

func (a *AnchovyTopping) getName() string {
	return a.pizza.getName() + " aux anchois"
}
