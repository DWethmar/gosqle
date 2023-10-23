package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

// SelectUsers selects users.
func SelectUsers(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
		clauses.Selectable{Expr: expressions.Column{Name: "name"}},
		clauses.Selectable{Expr: expressions.Column{Name: "email"}},
	).
		FromTable("addresses", nil).
		Limit(args.NewArgument(10)).
		Write(sb)
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
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, "", err
		}
		users = append(users, user)
	}
	return users, sb.String(), nil
}
