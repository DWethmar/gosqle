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

```

## Insert statement
Generate a insert query:
```go
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

```

## Delete statement
Generate a delete query:
```go
package main

import (
	"database/sql"
)

// Insert inserts a user.
func Delete(db *sql.DB) error {
	// TODO: Implement
	return nil
}

```

## Syntax used

![image](provision/images/SQL_syntax.svg)