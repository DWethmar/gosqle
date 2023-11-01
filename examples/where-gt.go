package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereGT selects users where id is greater than 10
func WhereGT(id int) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.Alias{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).
		Where(predicates.GT{
			Col:  expressions.Column{Name: "id"},
			Expr: args.NewArgument(id),
		}).Write(sb)
	if err != nil {
		return nil, "", err
	}
	return args.Args, sb.String(), nil
}
