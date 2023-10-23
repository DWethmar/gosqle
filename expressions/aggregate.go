package expressions

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/util"
)

var _ Expression = &Count{}

// Count is a count aggregate function.
type Count struct {
	Expr Expression
}

// WriteTo implements Expression.
func (c *Count) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "COUNT("); err != nil {
		return fmt.Errorf("error writing COUNT function: %v", err)
	}

	if err := c.Expr.Write(sw); err != nil {
		return fmt.Errorf("error writing COUNT expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing COUNT closing parenthesis: %v", err)
	}

	return nil
}

// NewCount returns a new Count aggregate functions expression.
func NewCount(column *Column) *Count {
	return &Count{
		Expr: column,
	}
}

var _ Expression = &Max{}

// Max is a max aggregate function.
type Max struct {
	Expr Expression
}

func (m *Max) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "MAX("); err != nil {
		return fmt.Errorf("error writing MAX function: %v", err)
	}

	if err := m.Expr.Write(sw); err != nil {
		return fmt.Errorf("error writing MAX expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing MAX closing parenthesis: %v", err)
	}

	return nil
}

// NewMax returns a new Max aggregate functions expression.
func NewMax(column *Column) *Max {
	return &Max{
		Expr: column,
	}
}

var _ Expression = &Min{}

// Min is a min aggregate function.
type Min struct {
	Expr Expression
}

func (m *Min) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "MIN("); err != nil {
		return fmt.Errorf("error writing MIN function: %v", err)
	}

	if err := m.Expr.Write(sw); err != nil {
		return fmt.Errorf("error writing MIN expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing MIN closing parenthesis: %v", err)
	}

	return nil
}

// NewMin returns a new Min aggregate functions expression.
func NewMin(column *Column) *Min {
	return &Min{
		Expr: column,
	}
}

var _ Expression = &Sum{}

// Sum is a sum aggregate function.
type Sum struct {
	Expr Expression
}

func (s *Sum) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "SUM("); err != nil {
		return fmt.Errorf("error writing SUM function: %v", err)
	}

	if err := s.Expr.Write(sw); err != nil {
		return fmt.Errorf("error writing SUM expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing SUM closing parenthesis: %v", err)
	}

	return nil
}

// NewSum returns a new Sum aggregate functions expression.
func NewSum(column *Column) *Sum {
	return &Sum{
		Expr: column,
	}
}

var _ Expression = &Avg{}

// Avg is a avg aggregate function.
type Avg struct {
	Expr Expression
}

func (a *Avg) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "AVG("); err != nil {
		return fmt.Errorf("error writing AVG function: %v", err)
	}

	if err := a.Expr.Write(sw); err != nil {
		return fmt.Errorf("error writing AVG expression: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing AVG closing parenthesis: %v", err)
	}

	return nil
}

// NewAvg returns a new Avg aggregate functions expression.
func NewAvg(column *Column) *Avg {
	return &Avg{
		Expr: column,
	}
}
