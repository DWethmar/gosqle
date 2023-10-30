package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses/set"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// UpdateUser updates a user.
func UpdateUser() (string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewUpdate("users").Set(set.Change{
		Col:  "name",
		Expr: args.NewArgument(fmt.Sprintf("new name %d", time.Now().Unix())),
	}).Where(predicates.EQ{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(1),
	}).Write(sb)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
