package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/postgres"
)

// InsertUser inserts a user.
func InsertUser() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewInsert("users", "name", "email").Values(
		args.NewArgument("John"),
		args.NewArgument(fmt.Sprintf("john%d@%s", time.Now().Unix(), "example.com")),
	).Write(sb)

	if err != nil {
		return nil, "", fmt.Errorf("error writing query: %v", err)
	}

	return args.Args, sb.String(), nil
}
