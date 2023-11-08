package examples

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
	).From(alias.NewStr("users")).
		Where(
			logic.And(predicates.Between(
				expressions.Column{Name: "id"},
				args.Create(10),
				args.Create(20),
			)),
			logic.AndGroup(
				logic.And(predicates.Between(
					expressions.Column{Name: "id"},
					args.Create(30),
					args.Create(40),
				)),
				logic.Or(predicates.EQ(
					expressions.Column{Name: "name"},
					args.Create("John"),
				)),
			),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values(), sb.String(), nil
}
