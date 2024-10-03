package main

type Expression interface {
	Interpret() bool
}
