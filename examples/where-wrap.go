package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereWrap selects users where id is between 10 and 20 or 30 and 40 or name is john
func WhereWrap(db *sql.DB) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.Between(
				expressions.Column{Name: "id"},
				args.NewArgument(10),
				args.NewArgument(20),
			)),
			logic.Or(
				logic.Group([]logic.Logic{
					logic.And(predicates.Between(
						expressions.Column{Name: "id"},
						args.NewArgument(30),
						args.NewArgument(40),
					)),
					logic.Or(predicates.EQ(
						expressions.Column{Name: "name"},
						args.NewArgument("John"),
					)),
				}),
			),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}
