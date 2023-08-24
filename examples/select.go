package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

type User struct {
	ID   int64
	Name string
}

// SelectUsers selects users.
func SelectUsers(db *sql.DB) ([]User, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	// SELECT u.id AS id, u.name AS name FROM users u LIMIT $1
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.NewColumn("id").SetFrom("u"), As: "id"},
		clauses.Selectable{Expr: expressions.NewColumn("name").SetFrom("u"), As: "name"},
	).From(expressions.Table{
		Name:  "users",
		Alias: "u",
	}).Limit(args.NewArgument(100)).WriteTo(sb)

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sb.String(), args.Args...)

	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
