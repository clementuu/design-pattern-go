package main

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (a *AndExpression) Interpret() bool {
	return a.expr1.Interpret() && a.expr2.Interpret()
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (o *OrExpression) Interpret() bool {
	return o.expr1.Interpret() || o.expr2.Interpret()
}
