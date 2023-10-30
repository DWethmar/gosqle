package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereBetween selects users where id is between 10 and 20
func WhereBetween(low, high int) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).
		Where(predicates.Between{
			Col:  expressions.Column{Name: "id"},
			Low:  args.NewArgument(low),
			High: args.NewArgument(high),
		}).Write(sb)
	if err != nil {
		return nil, "", err
	}
	return args.Args, sb.String(), nil
}
