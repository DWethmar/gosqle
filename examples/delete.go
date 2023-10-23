package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// Delete deletes a user.
func DeleteAddress(db *sql.DB) (string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	// DELETE FROM addresses WHERE user_id = $1
	err := gosqle.NewDelete("addresses").Where(predicates.EQ{
		Col:  expressions.Column{Name: "user_id"},
		Expr: args.NewArgument(111),
	}).Write(sb)
	if err != nil {
		return "", err
	}
	if _, err = db.Exec(sb.String(), args.Args...); err != nil {
		return "", err
	}

	return sb.String(), nil
}
