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

// WhereGTE selects users where id is greater than or equal to 10
func WhereGTE(id int) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.GTE(
				expressions.Column{Name: "id"},
				args.NewArgument(id),
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}
