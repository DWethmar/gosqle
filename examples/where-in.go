package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/from"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereIN selects users where name is in 'John', 'Jane' or 'Joe'.
// Example:
//
//	SELECT id FROM users WHERE name IN ($1, $2, $3);
func WhereIN(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).From(from.From{
		Expr: from.Table("users"),
	}).Where(predicates.In{
		Col: expressions.Column{Name: "id"},
		Expr: expressions.List{
			args.NewArgument("John"),
			args.NewArgument("Jane"),
			args.NewArgument("Joe"),
		},
	}).WriteTo(sb)

	if err != nil {
		return nil, "", err
	}

	rows, err := db.Query(sb.String(), args.Args...)
	if err != nil {
		return nil, "", err
	}

	var users []User
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID); err != nil {
			return nil, "", err
		}
		users = append(users, user)
	}

	return users, sb.String(), nil
}
