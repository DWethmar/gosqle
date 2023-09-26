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

type JoinType int

const (
	InnerJoin JoinType = iota
	LeftJoin
	RightJoin
	FullJoin
)

var joinTypes = map[JoinType]string{
	InnerJoin: "INNER JOIN",
	LeftJoin:  "LEFT JOIN",
	RightJoin: "RIGHT JOIN",
	FullJoin:  "FULL JOIN",
}

// JoinMatcher is a join clause.
type JoinMatcher interface {
	expressions.Expression
}

var _ JoinMatcher = (*JoinOn)(nil)

// JoinOn is a join clause that uses a list of predicates to match.
//
// Example:
//
//	LEFT JOIN table ON table.field = other_table.field AND table.field2 = other_table.field2
type JoinOn struct {
	Predicates []predicates.Predicate
}

func (j *JoinOn) WriteTo(sw io.StringWriter) error {
	if _, err := sw.WriteString("ON "); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	if err := predicates.WriteAll(sw, j.Predicates); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	return nil
}

var _ JoinMatcher = (*JoinUsing)(nil)

// JoinUsing is a join clause that uses a list of columns to match.
//
// Example:
//
//	LEFT JOIN table USING (field, field2)
type JoinUsing struct {
	Uses []string
}

func (j *JoinUsing) t() {}
func (j *JoinUsing) WriteTo(sw io.StringWriter) error {
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
func Write(sw io.StringWriter, joinType JoinType, from string, match JoinMatcher) error {
	if _, err := sw.WriteString(joinTypes[joinType]); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	if err := util.WriteStrings(sw, " ", from, " "); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	if err := match.WriteTo(sw); err != nil {
		return fmt.Errorf("failed to write JOIN: %w", err)
	}

	return nil
}

type Options struct {
	Type  JoinType
	From  string
	Match JoinMatcher // JoinOn or JoinUsing
}

// Clause represents a JOIN clause.
type Clause struct {
	joins []Options
}

func (j *Clause) Type() clauses.ClauseType { return clauses.JoinType }
func (j *Clause) WriteTo(sw io.StringWriter) error {
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
