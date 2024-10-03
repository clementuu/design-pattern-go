package main

type IShort interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Short struct {
	logo string
	size int
}

func (s *Short) setLogo(logo string) {
	s.logo = logo
}

func (s *Short) getLogo() string {
	return s.logo
}

func (s *Short) setSize(size int) {
	s.size = size
}

func (s *Short) getSize() int {
	return s.size
}
