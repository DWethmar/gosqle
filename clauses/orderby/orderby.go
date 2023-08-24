package orderby

import (
	"errors"
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var (
	_ clauses.Clause = &Clause{}
)

//go:generate stringer -type=Direction

// Direction of the order.
type Direction int

const (
	ASC Direction = iota
	DESC
)

// Column to order by.
type Sort struct {
	// Column to order by.
	*expressions.Column
	// Direction of the order.
	Direction Direction
}

// NewSort creates a new Sort.
func NewSort(column *expressions.Column, direction Direction) Sort {
	return Sort{
		Column:    column,
		Direction: direction,
	}
}

// WriteTo writes a ORDER BY clause to the given string writer.
func (s Sort) WriteTo(sw io.StringWriter) error {
	if err := s.Column.WriteTo(sw); err != nil {
		return fmt.Errorf("failed to write ORDER BY column: %v", err)
	}

	if err := util.WriteStrings(sw, " ", s.Direction.String()); err != nil {
		return fmt.Errorf("failed to write ORDER BY direction: %v", err)
	}

	return nil
}

// Write writes a ORDER BY clause to the given string writer.
// Add direction to the column name to change the direction.
//
// Example:
//
//	Mysql:
//		ORDER BY field1 DESC, field2 ASC, field3 ASC
//	Postgres:
//		ORDER BY field1 DESC, field2 ASC, field3 ASC
func Write(sw io.StringWriter, sorting []Sort) error {
	if len(sorting) == 0 {
		return errors.New("no order given")
	}

	if _, err := sw.WriteString(`ORDER BY `); err != nil {
		return fmt.Errorf("failed to write ORDER BY: %v", err)
	}

	for i, o := range sorting {
		if err := o.WriteTo(sw); err != nil {
			return fmt.Errorf("failed to write ORDER BY column: %v", err)
		}

		if i < len(sorting)-1 {
			if _, err := sw.WriteString(", "); err != nil {
				return fmt.Errorf("failed to write ORDER BY separator: %v", err)
			}
		}
	}

	return nil
}

// Clause represents a ORDER BY clause.
type Clause struct {
	sorting []Sort
}

func (o *Clause) Type() clauses.ClauseType         { return clauses.OrderByType }
func (o *Clause) WriteTo(sw io.StringWriter) error { return Write(sw, o.sorting) }

func New(sorting []Sort) *Clause {
	return &Clause{
		sorting: sorting,
	}
}
