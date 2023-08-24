package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/postgres"
)

// InsertUser inserts a user.
func InsertUser(db *sql.DB) (string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	// INSERT INTO users (name, email) VALUES ($1, $2)
	err := gosqle.NewInsert("users", "name", "email").Values(
		args.NewArgument("John"),
		args.NewArgument(fmt.Sprintf("john%d@%s", time.Now().Unix(), "example.com")),
	).WriteTo(sb)

	if err != nil {
		return "", err
	}

	if _, err = db.Exec(sb.String(), args.Args...); err != nil {
		return "", err
	}

	return sb.String(), nil
}
