package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLike selects users where name is like anna%
func WhereLike() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.Alias{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).
		Where(predicates.Like{
			Col:  expressions.Column{Name: "name"},
			Expr: args.NewArgument("anna%"),
		}).Write(sb)
	if err != nil {
		return nil, "", err
	}
	return args.Args, sb.String(), nil
}
