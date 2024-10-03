package main

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 3
}

func (c *CheeseTopping) getName() string {
	return c.pizza.getName() + " au fromage"
}
