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

// WhereWrap selects users where id is between 10 and 20 or 30 and 40 or name is john
// Example:
//
//	SELECT id FROM users WHERE (id BETWEEN $1 AND $2 OR id BETWEEN $3 AND $4) OR name = $5;
func WhereWrap(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).From(from.From{
		Expr: from.Table("users"),
	}).Where(
		predicates.Wrap{
			Predicates: []predicates.Predicate{
				predicates.Between{
					Col:  expressions.Column{Name: "id"},
					Low:  args.NewArgument(10),
					High: args.NewArgument(20),
				},
				predicates.Between{
					Col:  expressions.Column{Name: "id"},
					Low:  args.NewArgument(30),
					High: args.NewArgument(40),
				},
			},
		}, predicates.EQ{
			Col:   expressions.Column{Name: "name"},
			Expr:  args.NewArgument("John"),
			Logic: predicates.OR,
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
		if err = rows.Scan(&user.ID); err != nil {
			return nil, "", err
		}
		users = append(users, user)
	}

	return users, sb.String(), nil
}