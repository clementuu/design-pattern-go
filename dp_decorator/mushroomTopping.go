package main

type MushroomTopping struct {
	pizza IPizza
}

func (m *MushroomTopping) getPrice() int {
	return m.pizza.getPrice() + 2
}

func (m *MushroomTopping) getName() string {
	return m.pizza.getName() + " aux champignons"
}
