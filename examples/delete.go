package examples

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// Delete deletes a user.
func DeleteAddress(id int64) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewDelete("addresses").
		Where(
			logic.And(predicates.EQ(
				expressions.Column{Name: "id", From: "addresses"},
				args.Create(id),
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values(), sb.String(), nil
}
