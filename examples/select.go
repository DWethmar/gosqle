package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

// SelectUsers selects users.
func SelectUsers() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
		clauses.Selectable{Expr: expressions.Column{Name: "name"}},
		clauses.Selectable{Expr: expressions.Column{Name: "email"}},
	).
		FromTable("users", nil).
		Limit(args.NewArgument(10)).
		Write(sb)
	if err != nil {
		return nil, "", err
	}
	return args.Args, sb.String(), nil
}
