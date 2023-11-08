package examples

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

// SelectUsers selects users.
func SelectUsers(limit int64) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
		alias.New(expressions.Column{Name: "name"}),
		alias.New(expressions.Column{Name: "email"}),
	).
		From(alias.NewStr("users")).
		Limit(args.Create(limit)).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values(), sb.String(), nil
}
