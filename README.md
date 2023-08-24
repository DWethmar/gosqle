# gosqle
gosqle is a golang package that can generate sql queries. 

Table of Contents:
- [gosqle](#gosqle)
  - [Examples](#examples)
    - [Select statement](#select-statement)
      - [Generate a select query:](#generate-a-select-query)
      - [Generate select query using group by and aggregate functions:](#generate-select-query-using-group-by-and-aggregate-functions)
    - [Insert statement](#insert-statement)
      - [Generate a insert query:](#generate-a-insert-query)
    - [Delete statement](#delete-statement)
      - [Generate a delete query:](#generate-a-delete-query)
    - [Update statement](#update-statement)
      - [Generate a update query:](#generate-a-update-query)
  - [Syntax used](#syntax-used)

## Examples
Examples shown here are generated into this README.md file from the [examples](examples) folder. See [README.tmpl.md](README.tmpl.md) for more information.
To generate the examples into this README.md file, run the following command:
```bash
./run_readme.sh
```

To run examples from the [examples](examples) folder, run the following command:
```bash
docker-compose up -d
./run_examples.sh
```

### Select statement
Create a select statement with the following syntax:
```go
gosqle.NewSelect(...columns)
```
#### Generate a select query:
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

	// SELECT u.id AS id, u.name AS name, u.email AS email FROM users u LIMIT $1
	err := gosqle.NewSelect(
		clauses.Selectable{
			Expr: expressions.NewColumn("id").SetFrom("u"),
			As:   "id",
		},
		clauses.Selectable{
			Expr: expressions.NewColumn("name").SetFrom("u"),
			As:   "name",
		},
		clauses.Selectable{
			Expr: expressions.NewColumn("email").SetFrom("u"),
			As:   "email",
		},
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

#### Generate select query using group by and aggregate functions:
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/groupby"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

type AmountOfAddressesPerCountry struct {
	Country string
	Count   int64
}

// SelectAmountOfAddressesPerCountry select amount of addresses per country
func SelectAmountOfAddressesPerCountry(db *sql.DB) ([]AmountOfAddressesPerCountry, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	/**
	SELECT country, COUNT(id) AS address_count
	FROM addresses
	GROUP BY country
	ORDER BY address_count DESC;
	**/
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.NewColumn("country")},
		clauses.Selectable{Expr: expressions.NewCount(expressions.NewColumn("id")), As: "address_count"},
	).From(expressions.Table{
		Name: "addresses",
	}).GroupBy(groupby.ColumnGrouping{
		expressions.NewColumn("country"),
	}).OrderBy(orderby.Sort{
		Column:    expressions.NewColumn("address_count"),
		Direction: orderby.DESC,
	}).WriteTo(sb)
	if err != nil {
		return nil, "", err
	}

	rows, err := db.Query(sb.String(), args.Args...)
	if err != nil {
		return nil, "", err
	}

	var r []AmountOfAddressesPerCountry
	for rows.Next() {
		var a AmountOfAddressesPerCountry
		err = rows.Scan(&a.Country, &a.Count)
		if err != nil {
			return nil, "", err
		}
		r = append(r, a)
	}

	return r, sb.String(), nil
}

```

### Insert statement
```go
gosqle.NewInsert(table, ...columns)
```
#### Generate a insert query:
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

### Delete statement
#### Generate a delete query:
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

### Update statement
#### Generate a update query:
```go
package main

import (
	"database/sql"
)

// Update updates a user.
func Update(db *sql.DB) error {
	// TODO: Implement
	return nil
}

```

## Syntax used

![image](provision/images/SQL_syntax.svg)