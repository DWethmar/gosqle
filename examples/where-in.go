package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereIN selects users where name is in names.
func WhereIN(names []string) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	list := []expressions.Expression{}
	for _, name := range names {
		list = append(list, args.NewArgument(name))
	}
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).
		Where(predicates.In{
			Col:  expressions.Column{Name: "id"},
			Expr: expressions.List(list),
		}).Write(sb)
	if err != nil {
		return nil, "", err
	}
	return args.Args, sb.String(), nil
}
