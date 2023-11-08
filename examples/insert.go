package examples

import (
	"fmt"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/postgres"
)

// InsertUser inserts a user.
func InsertUser(name, email string) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewInsert("users", "name", "email").Values(
		args.Create(name),
		args.Create(email),
	).Write(sb)

	if err != nil {
		return nil, "", fmt.Errorf("error writing query: %v", err)
	}

	return args.Values(), sb.String(), nil
}
