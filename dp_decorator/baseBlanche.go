package main

type BaseBlanche struct {
}

func (p *BaseBlanche) getPrice() int {
	return 6
}

func (p *BaseBlanche) getName() string {
	return "Blanche"
}
