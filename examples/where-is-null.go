package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereIsNull selects addresses where phone is null
func WhereIsNull() (string, error) {
	sb := new(strings.Builder)
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).
		Where(predicates.IsNull{
			Col: expressions.Column{Name: "phone"},
		}).Write(sb)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
