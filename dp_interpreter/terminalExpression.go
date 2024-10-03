package main

import "strings"

type TerminalExpression struct {
	data string
}

func (t *TerminalExpression) Interpret() bool {
	return strings.Contains(t.data, "hello")
}
