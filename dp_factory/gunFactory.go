package main

import "fmt"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	if gunType == "gatling" {
		return newGatling(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
