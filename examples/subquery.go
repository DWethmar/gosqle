package main

import (
	"fmt"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// SelectUsers selects users.
func PeopleOfAmsterdam() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "name"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.IN(
				expressions.Column{Name: "id"},
				gosqle.NewSelect(
					alias.New(expressions.Column{Name: "user_id"}),
				).
					FromTable("addresses", nil).
					Where(
						logic.And(predicates.EQ(
							expressions.Column{Name: "city"},
							args.NewArgument("Amsterdam"),
						)),
					).Statement, // <-- This is the sub-query without semicolon
			)),
		).Write(sb)

	if err != nil {
		return nil, "", fmt.Errorf("error writing query: %v", err)
	}

	return args.Values, sb.String(), nil
}
