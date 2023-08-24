# gosqle
gosqle is a golang package that can generate sql queries. 

Table of Contents:
- [gosqle](#gosqle)
  - [Select statement](#select-statement)
  - [Insert statement](#insert-statement)
  - [Delete statement](#delete-statement)
  - [Syntax used](#syntax-used)

## Select statement
Generate a select query:
```go
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
	ID    int64
	Name  string
	Email string
}

// SelectUsers selects users.
func SelectUsers(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	// SELECT u.id AS id, u.name AS name FROM users u LIMIT $1
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.NewColumn("id").SetFrom("u"), As: "id"},
		clauses.Selectable{Expr: expressions.NewColumn("name").SetFrom("u"), As: "name"},
		clauses.Selectable{Expr: expressions.NewColumn("email").SetFrom("u"), As: "email"},
	).From(expressions.Table{
		Name:  "users",
		Alias: "u",
	}).Limit(args.NewArgument(10)).WriteTo(sb)
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

```

## Insert statement
Generate a insert query:
```go
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

```

## Delete statement
Generate a delete query:
```go
package main

import (
	"database/sql"
)

// Delete deletes a user.
func Delete(db *sql.DB) error {
	// TODO: Implement
	return nil
}

```

## Syntax used

![image](provision/images/SQL_syntax.svg)