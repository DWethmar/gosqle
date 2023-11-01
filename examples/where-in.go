package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
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
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.IN(
				expressions.Column{Name: "name"},
				list...,
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}
