package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLTE selects users where id is less than or equal to 10
func WhereLTE() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.Alias{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.LT{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(10),
	}).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Args, sb.String(), nil
}
