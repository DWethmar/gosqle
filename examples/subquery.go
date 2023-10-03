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

// SelectUsers selects users.
func PeopleOfAmsterdam(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	// SELECT name
	// FROM users
	// WHERE id IN (
	//     SELECT user_id
	//     FROM addresses
	//     WHERE city = 'New York'
	// );
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "name"}},
	).From(from.From{
		Expr: from.Table("users"),
	}).Where(
		predicates.In{
			Col: expressions.Column{Name: "id"},
			Expr: gosqle.NewSelect(
				clauses.Selectable{Expr: expressions.Column{Name: "user_id"}},
			).From(from.From{
				Expr: from.Table("addresses"),
			}).Where(
				predicates.EQ{
					Col:  expressions.Column{Name: "city"},
					Expr: args.NewArgument("Amsterdam"),
				},
			).Statement, // <- This is the subquery, so without semicolon.
		},
	).WriteTo(sb)

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
		err = rows.Scan(&user.Name)
		if err != nil {
			return nil, "", err
		}
		users = append(users, user)
	}

	return users, sb.String(), nil
}
