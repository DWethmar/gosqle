# gosqle
gosqle is a golang package that can generate sql queries. 

Table of Contents:
- [gosqle](#gosqle)
  - [Examples](#examples)
    - [Select](#select)
      - [Generate a select query:](#generate-a-select-query)
      - [Generate select query using group by and aggregate functions:](#generate-select-query-using-group-by-and-aggregate-functions)
      - [Subquery](#subquery)
    - [Insert](#insert)
      - [Generate an insert query:](#generate-an-insert-query)
    - [Delete](#delete)
      - [Generate a delete query:](#generate-a-delete-query)
    - [Update](#update)
      - [Generate an update query:](#generate-an-update-query)
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
# if needed you can seed the database with some data
./run_seed.sh
./run_examples.sh
```

### Select
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
	"github.com/dwethmar/gosqle/clauses/from"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

// SelectUsers selects users.
func SelectUsers(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	// SELECT id, name, email FROM users LIMIT 10;
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
		clauses.Selectable{Expr: expressions.Column{Name: "name"}},
		clauses.Selectable{Expr: expressions.Column{Name: "email"}},
	).From(from.From{
		Expr: from.Table("users"),
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
	"github.com/dwethmar/gosqle/clauses/from"
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
		clauses.Selectable{
			Expr: &expressions.Column{Name: "country"},
		},
		clauses.Selectable{
			Expr: expressions.NewCount(&expressions.Column{Name: "id"}),
			As:   "address_count",
		},
	).From(from.From{
		Expr: from.Table("addresses"),
	}).GroupBy(groupby.ColumnGrouping{
		&expressions.Column{Name: "country"},
	}).OrderBy(orderby.Sort{
		Column:    &expressions.Column{Name: "address_count"},
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

#### Subquery
```go
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

```

### Insert
```go
gosqle.NewInsert(table, ...columns)
```
#### Generate an insert query:
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

### Delete
#### Generate a delete query:
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// Delete deletes a user.
func DeleteAddress(db *sql.DB) (string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	// DELETE FROM addresses WHERE user_id = $1
	err := gosqle.NewDelete("addresses").Where(
		predicates.EQ{
			Col:  expressions.Column{Name: "user_id"},
			Expr: args.NewArgument(111),
		},
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

### Update
#### Generate an update query:
```go
package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses/set"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// UpdateUser updates a user.
func UpdateUser(db *sql.DB) (string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	// UPDATE users SET name = $1 WHERE id = $2
	err := gosqle.NewUpdate("users").Set(set.Change{
		Col:  "name",
		Expr: args.NewArgument(fmt.Sprintf("new name %d", time.Now().Unix())),
	}).Where(predicates.EQ{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(1),
	}).WriteTo(sb)
	if err != nil {
		return "", err
	}
	if _, err = db.Exec(sb.String(), args.Args...); err != nil {
		return "", err
	}
	return sb.String(), nil
}

```

## Syntax used

![image](provision/images/SQL_syntax.svg)