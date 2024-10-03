package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
	makeShort() IShort
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
