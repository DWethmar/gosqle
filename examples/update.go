package examples

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses/set"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// UpdateUser updates a user.
func UpdateUser(userID int64, name string) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewUpdate("users").Set(set.Change{
		Col:  "name",
		Expr: args.Create(name),
	}).Where(
		logic.And(predicates.EQ(
			expressions.Column{Name: "id"},
			args.Create(userID),
		)),
	).Write(sb)
	if err != nil {
		return nil, "", err
	}
	return args.Values(), sb.String(), nil
}
