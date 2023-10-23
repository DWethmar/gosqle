package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLike selects users where name is like anna%
func WhereLike(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.Like{
		Col:  expressions.Column{Name: "name"},
		Expr: args.NewArgument("anna%"),
	}).Write(sb)

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
