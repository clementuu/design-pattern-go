package main

type VegetableTopping struct {
	pizza IPizza
}

func (v *VegetableTopping) getPrice() int {
	return v.pizza.getPrice() + 4
}

func (v *VegetableTopping) getName() string {
	return v.pizza.getName() + " aux légumes"
}
