package main

type Gatling struct {
	Gun
}

func newGatling() IGun {
	return &Gatling{
		Gun: Gun{
			name:  "Gatling gun",
			power: 35,
		},
	}
}
