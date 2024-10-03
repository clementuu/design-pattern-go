package main

type HamTopping struct {
	pizza IPizza
}

func (h *HamTopping) getPrice() int {
	return h.pizza.getPrice() + 2
}

func (h *HamTopping) getName() string {
	return h.pizza.getName() + " au jambon"
}
