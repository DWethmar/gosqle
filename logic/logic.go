package logic

import (
	"errors"
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
)

type Operator int

const (
	AndOperator Operator = iota
	OrOperator
)

type Logic struct {
	Operator  Operator
	Condition expressions.Expression
}

func (l *Logic) Write(sw io.StringWriter) error {
	if l.Condition == nil {
		return errors.New("logic condition is nil")
	}

	switch l.Operator {
	case AndOperator:
		if _, err := sw.WriteString("AND "); err != nil {
			return fmt.Errorf("error writing AND: %v", err)
		}
	case OrOperator:
		if _, err := sw.WriteString("OR "); err != nil {
			return fmt.Errorf("error writing OR: %v", err)
		}
	default:
		return fmt.Errorf("unknown logic operator: %v", l.Operator)
	}

	if err := l.Condition.Write(sw); err != nil {
		return fmt.Errorf("error writing logic condition: %v", err)
	}

	return nil
}

type Group []Logic

func (c Group) Write(sw io.StringWriter) error {
	if len(c) == 0 {
		return errors.New("Group is empty")
	}

	return Where(sw, c)
}

// And creates logic with AND operator
func And(condition expressions.Expression) Logic {
	return Logic{
		Operator:  AndOperator,
		Condition: condition,
	}
}

// Or creates logic with OR operator
func Or(condition expressions.Expression) Logic {
	return Logic{
		Operator:  OrOperator,
		Condition: condition,
	}
}
