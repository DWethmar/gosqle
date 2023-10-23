package gosqle

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/from"
	"github.com/dwethmar/gosqle/clauses/groupby"
	"github.com/dwethmar/gosqle/clauses/having"
	"github.com/dwethmar/gosqle/clauses/join"
	"github.com/dwethmar/gosqle/clauses/limit"
	"github.com/dwethmar/gosqle/clauses/offset"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/clauses/where"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/predicates"
	"github.com/dwethmar/gosqle/statement"
)

// Select is a wrapper for a select query statement.
type Select struct {
	// Statement is the select statement.
	statement.Statement
}

func (s *Select) From(f from.From) *Select {
	return s.SetClause(from.New(f))
}

type FromTableOptions struct {
	Alias string
}

func (s *Select) FromTable(table string, opt *FromTableOptions) *Select {
	var as string
	if opt != nil {
		as = opt.Alias
	}

	return s.SetClause(
		from.New(from.NewFrom(table, as)),
	)
}

// Join adds a join clause to the select statement.
// If no join options are given, the join clause will be ignored.
func (s *Select) Join(j ...join.Options) *Select {
	if len(j) == 0 {
		return s
	}

	return s.SetClause(join.New(j))
}

// Where adds a where clause to the select statement.
// If no predicates are given, the where clause will be ignored.
func (s *Select) Where(predicates ...predicates.Predicate) *Select {
	if len(predicates) == 0 {
		return s
	}

	return s.SetClause(where.New(predicates))
}

// GroupBy adds a group by clause to the select statement.
//
// Grouping options:
// - groupby.ColumnGrouping
// - <add more>
func (s *Select) GroupBy(grouping groupby.Grouping) *Select {
	return s.SetClause(groupby.New(grouping))
}

// Having adds a having clause to the select statement.
func (s *Select) Having(p ...predicates.Predicate) *Select {
	return s.SetClause(having.New(p))
}

// OrderBy adds a order by clause to the select statement.
// If no sorting options are given, the order by clause will be ignored.
func (s *Select) OrderBy(sorting ...orderby.Sort) *Select {
	if len(sorting) == 0 {
		return s
	}

	return s.SetClause(orderby.New(sorting))
}

// Limit adds a limit clause to the select statement.
func (s *Select) Limit(argument expressions.Expression) *Select {
	return s.SetClause(limit.New(argument))
}

// Offset adds a offset clause to the select statement.
func (s *Select) Offset(argument expressions.Expression) *Select {
	return s.SetClause(offset.New(argument))
}

// SetClause sets the clause for the select statement.
func (s *Select) SetClause(c clauses.Clause) *Select {
	s.Statement.SetClause(c)
	return s
}

// Write writes the select query to the given writer.
// It also adds a semicolon to the end of the query.
func (s *Select) Write(sw io.StringWriter) error {
	if err := s.Statement.Write(sw); err != nil {
		return fmt.Errorf("failed to write select statement: %v", err)
	}

	// Add a semicolon to the end of the query.
	if _, err := sw.WriteString(";"); err != nil {
		return fmt.Errorf("failed to write semicolon: %v", err)
	}

	return nil
}

// NewSelect creates a new select query.
func NewSelect(
	selectables ...clauses.Selectable,
) *Select {
	return &Select{
		Statement: statement.NewSelect(selectables),
	}
}
