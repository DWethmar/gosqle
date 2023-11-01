package predicates

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var _ expressions.Expression = &Comparison{}

// Comparison is predicate that evaluates to TRUE, FALSE, or UNKNOWN
type Comparison struct {
	Left     expressions.Expression
	Operator string
	Right    expressions.Expression
}

func (c *Comparison) Write(writer io.StringWriter) error {
	if writer == nil {
		return fmt.Errorf("writer is nil")
	}

	if c.Left == nil {
		return fmt.Errorf("comparison left is nil")
	}

	if c.Operator == "" {
		return fmt.Errorf("comparison operator is empty")
	}

	if c.Right == nil {
		return fmt.Errorf("comparison right is nil")
	}

	if err := c.Left.Write(writer); err != nil {
		return fmt.Errorf("error writing comparison left: %v", err)
	}

	if err := util.WriteStrings(writer, " ", c.Operator, " "); err != nil {
		return fmt.Errorf("error writing comparison operator: %v", err)
	}

	if err := c.Right.Write(writer); err != nil {
		return fmt.Errorf("error writing comparison right: %v", err)
	}

	return nil
}

// EQ creates a new comparison predicate
func EQ(left expressions.Expression, right expressions.Expression) *Comparison {
	return &Comparison{
		Left:     left,
		Operator: "=",
		Right:    right,
	}
}

// NE creates a new comparison predicate
func NE(left expressions.Expression, right expressions.Expression) *Comparison {
	return &Comparison{
		Left:     left,
		Operator: "!=",
		Right:    right,
	}
}

func GT(left expressions.Expression, right expressions.Expression) *Comparison {
	return &Comparison{
		Left:     left,
		Operator: ">",
		Right:    right,
	}
}

func GTE(left expressions.Expression, right expressions.Expression) *Comparison {
	return &Comparison{
		Left:     left,
		Operator: ">=",
		Right:    right,
	}
}

func LT(left expressions.Expression, right expressions.Expression) *Comparison {
	return &Comparison{
		Left:     left,
		Operator: "<",
		Right:    right,
	}
}

func LTE(left expressions.Expression, right expressions.Expression) *Comparison {
	return &Comparison{
		Left:     left,
		Operator: "<=",
		Right:    right,
	}
}

func Like(left expressions.Expression, right expressions.Expression) *Comparison {
	return &Comparison{
		Left:     left,
		Operator: "LIKE",
		Right:    right,
	}
}

func Not(p *Comparison) expressions.ExpressionFunc {
	return expressions.ExpressionFunc(func(sw io.StringWriter) error {
		if err := expressions.Prepend(sw, "NOT ", p); err != nil {
			return fmt.Errorf("failed to write NOT")
		}

		return nil
	})
}
