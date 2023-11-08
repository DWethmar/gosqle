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

// WhereNOT selects users where name is not John.
func WhereNOT(db *sql.DB) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).From(alias.NewStr("users")).
		Where(
			logic.And(predicates.Not(predicates.EQ(
				expressions.Column{Name: "name"},
				args.Create("John"),
			))),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values(), sb.String(), nil
}
