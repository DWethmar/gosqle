package predicates

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
)

// In creates a new in predicate.
func IN(target expressions.Expression, expr ...expressions.Expression) *Comparison {
	return &Comparison{
		Left:     target,
		Operator: "IN",
		Right: expressions.ExpressionFunc(func(writer io.StringWriter) error {
			if err := expressions.WrapInParenthesis(writer, expressions.List(expr)); err != nil {
				return fmt.Errorf("error wrapping IN expression in parenthesis: %v", err)
			}

			return nil
		}),
	}
}
