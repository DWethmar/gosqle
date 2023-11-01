package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses/set"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// UpdateUser updates a user.
func UpdateUser() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewUpdate("users").Set(set.Change{
		Col:  "name",
		Expr: args.NewArgument(fmt.Sprintf("new name %d", time.Now().Unix())),
	}).Where(
		logic.And(predicates.EQ(
			expressions.Column{Name: "id"},
			args.NewArgument(193),
		)),
	).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}
