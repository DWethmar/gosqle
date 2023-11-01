package main

import (
	"fmt"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// SelectUsers selects users.
func PeopleOfAmsterdam() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.Alias{Expr: expressions.Column{Name: "name"}},
	).FromTable("users", nil).
		Where(predicates.In{
			Col: expressions.Column{Name: "id"},
			Expr: gosqle.NewSelect(
				alias.Alias{Expr: expressions.Column{Name: "user_id"}},
			).FromTable("addresses", nil).Where(predicates.EQ{
				Col:  expressions.Column{Name: "city"},
				Expr: args.NewArgument("Amsterdam"),
			}).Statement, // <- This is the subquery, so without semicolon.
		}).Write(sb)

	if err != nil {
		return nil, "", fmt.Errorf("error writing query: %v", err)
	}

	return args.Args, sb.String(), nil
}
