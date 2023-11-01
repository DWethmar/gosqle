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

// WhereNE selects users where name is not equal to 'John'.
func WhereNE() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.NE(
				expressions.Column{Name: "name"},
				args.NewArgument("John"),
			)),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}
