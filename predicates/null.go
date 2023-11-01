package predicates

import (
	"github.com/dwethmar/gosqle/expressions"
)

func IsNull(target expressions.Expression) *Comparison {
	return &Comparison{
		Left:     target,
		Operator: "IS",
		Right:    expressions.String("NULL"),
	}
}

func IsNotNull(target expressions.Expression) *Comparison {
	return &Comparison{
		Left:     target,
		Operator: "IS NOT",
		Right:    expressions.String("NULL"),
	}
}
