package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/postgres"
)

// Insert inserts a user.
func Insert(db *sql.DB) error {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	// INSERT INTO users (id, name) VALUES ($1, $2)
	err := gosqle.NewInsert("users",
		"id",
		"name",
	).Values(
		args.NewArgument(1),
		args.NewArgument("John"),
	).WriteTo(sb)

	if err != nil {
		return err
	}

	if _, err = db.Exec(sb.String(), args.Args...); err != nil {
		return err
	}

	return nil
}
