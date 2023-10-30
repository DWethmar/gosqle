package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLT selects users where id is less than 10
func WhereLT() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.LT{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(10),
	}).Write(sb)
	if err != nil {
		return nil, "", err
	}
	return args.Args, sb.String(), nil
}
