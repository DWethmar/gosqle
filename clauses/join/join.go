package join

import (
	"fmt"
	"io"
	"strings"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/predicates"
	"github.com/dwethmar/gosqle/util"
)

type Type int

const (
	InnerJoin Type = iota
	LeftJoin
	RightJoin
	FullJoin
)

var types = map[Type]string{
	InnerJoin: "INNER JOIN",
	LeftJoin:  "LEFT JOIN",
	RightJoin: "RIGHT JOIN",
	FullJoin:  "FULL JOIN",
}

// Matcher is the interface that wraps the matching part of a JOIN clause.
// It can be either an On or Using clause.
type Matcher interface {
	expressions.Expression
}

var _ Matcher = (*On)(nil)

// On is a join clause that uses a list of predicates to match.
//
// Example:
//
//	LEFT JOIN table ON table.field = other_table.field AND table.field2 = other_table.field2
type On struct {
	Predicates []predicates.Predicate
}

func (j *On) Write(sw io.StringWriter) error {
	if _, err := sw.WriteString("ON "); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	if err := predicates.WriteAll(sw, j.Predicates); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	return nil
}

var _ Matcher = (*Using)(nil)

// Using is a join clause that uses a list of columns to match.
//
// Example:
//
//	LEFT JOIN table USING (field, field2)
type Using struct {
	Uses []string
}

func (j *Using) t() {}
func (j *Using) Write(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "USING (", strings.Join(j.Uses, ", "), ")"); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	return nil
}

// Join writes a JOIN clause to a strings.Builder.
//
// Example:
//
//	LEFT JOIN table ON table.field = other_table.field AND table.field2 = other_table.field2
//	LEFT JOIN table USING (field, field2)
func Write(sw io.StringWriter, joinType Type, from string, match Matcher) error {
	if _, err := sw.WriteString(types[joinType]); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	if err := util.WriteStrings(sw, " ", from, " "); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	if err := match.Write(sw); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	return nil
}

type Options struct {
	Type  Type
	From  string
	Match Matcher // On or Using
}

// Clause represents a JOIN clause.
type Clause struct {
	joins []Options
}

func (j *Clause) Type() clauses.ClauseType { return clauses.JoinType }
func (j *Clause) Write(sw io.StringWriter) error {
	for i, join := range j.joins {
		if err := Write(sw, join.Type, join.From, join.Match); err != nil {
			return fmt.Errorf("failed to write JOIN: %w", err)
		}

		if i < len(j.joins)-1 {
			if _, err := sw.WriteString(" "); err != nil {
				return fmt.Errorf("failed to write JOIN: %w", err)
			}
		}
	}

	return nil
}

// New creates a new join clause.
func New(j []Options) *Clause {
	return &Clause{
		joins: j,
	}
}
