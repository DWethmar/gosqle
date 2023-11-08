package examples

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereGT selects users where id is greater than 10
func WhereGT(id int) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).From(alias.NewStr("users")).
		Where(
			logic.And(predicates.GT(
				expressions.Column{Name: "id"},
				args.Create(id),
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values(), sb.String(), nil
}
