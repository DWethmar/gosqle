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

// WhereWrap selects users where id is between 10 and 20 or 30 and 40 or name is john
func WhereWrap(db *sql.DB) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.Alias{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).
		Where(
			predicates.Wrap{
				Predicates: []predicates.Predicate{
					predicates.Between{
						Col:  expressions.Column{Name: "id"},
						Low:  args.NewArgument(10),
						High: args.NewArgument(20),
					},
					predicates.Between{
						Col:  expressions.Column{Name: "id"},
						Low:  args.NewArgument(30),
						High: args.NewArgument(40),
					},
				},
			},
			predicates.EQ{
				Col:   expressions.Column{Name: "name"},
				Expr:  args.NewArgument("John"),
				Logic: predicates.OR,
			},
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Args, sb.String(), nil
}
