package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// Delete deletes a user.
func DeleteAddress() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewDelete("addresses").Where(predicates.EQ{
		Col:  expressions.Column{Name: "user_id"},
		Expr: args.NewArgument(111),
	}).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Args, sb.String(), nil
}
