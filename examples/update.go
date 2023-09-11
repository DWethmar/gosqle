package main

import (
	"database/sql"
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
func UpdateUser(db *sql.DB) (string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	// UPDATE users SET name = $1 WHERE id = $2
	err := gosqle.NewUpdate("users").Set(set.Change{
		Col:  "name",
		Expr: args.NewArgument(fmt.Sprintf("new name %d", time.Now().Unix())),
	}).Where(predicates.EQ{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(1),
	}).WriteTo(sb)
	if err != nil {
		return "", err
	}
	if _, err = db.Exec(sb.String(), args.Args...); err != nil {
		return "", err
	}
	return sb.String(), nil
}
