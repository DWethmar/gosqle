package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereNOT selects users where name is not John.
func WhereNOT(db *sql.DB) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.Alias{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).
		Where(predicates.Not{
			Predicate: predicates.EQ{
				Col:  expressions.Column{Name: "name"},
				Expr: args.NewArgument("John"),
			},
		}).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Args, sb.String(), nil
}
