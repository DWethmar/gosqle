package groupby

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
)

// WriteColumns writes a group by clause to the given writer.
func WriteGroupByColumns(sw io.StringWriter, columns []*expressions.Column) error {
	if _, err := sw.WriteString("GROUP BY "); err != nil {
		return fmt.Errorf("unable to write GROUP BY: %v", err)
	}
	for i, column := range columns {
		if err := column.WriteTo(sw); err != nil {
			return fmt.Errorf("unable to write GROUP BY column at %d: %v", i, err)
		}

		if i < len(columns)-1 {
			if _, err := sw.WriteString(", "); err != nil {
				return fmt.Errorf("unable to write GROUP BY separator at %d: %v", i, err)
			}
		}
	}

	return nil
}

// Grouping represents how to group by.
type Grouping interface {
	expressions.Expression
	g()
}

var _ Grouping = &ColumnGrouping{}

// ColumnGrouping represents a group by clause with columns
type ColumnGrouping struct {
	Columns []*expressions.Column
}

func (c *ColumnGrouping) WriteTo(sw io.StringWriter) error { return WriteGroupByColumns(sw, c.Columns) }
func (*ColumnGrouping) g()                                 {}

var (
	_ clauses.Clause = &Clause{}
)

// Clause represents a group by clause.
type Clause struct {
	grouping Grouping
}

func (*Clause) Type() clauses.ClauseType { return clauses.GroupByType }
func (c *Clause) WriteTo(sw io.StringWriter) error {
	if err := c.grouping.WriteTo(sw); err != nil {
		return fmt.Errorf("unable to write GROUP BY columns: %v", err)
	}

	return nil
}

// New creates a new group by clause.
func New(grouping Grouping) *Clause {
	return &Clause{
		grouping: grouping,
	}
}
