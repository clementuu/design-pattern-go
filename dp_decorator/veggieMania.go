package main

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 8
}

func (p *VeggieMania) getName() string {
	return "VeggieMania"
}
