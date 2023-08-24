package clauses

import (
	"github.com/dwethmar/gosqle/expressions"
)

//go:generate stringer -type=ClauseType

type ClauseType int

const (
	UnknownType ClauseType = iota
	FromType
	JoinType
	WhereType
	LimitType
	OffsetType
	OrderByType
	ReturningType
	ValuesType
	SetType
	GroupByType
	HavingType
)

// CLause
type Clause interface {
	expressions.Expression
	Type() ClauseType
}
