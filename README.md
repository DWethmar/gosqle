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
    - [Where conditions](#where-conditions)
      - [equal](#equal)
      - [Not equal](#not-equal)
      - [Greater than](#greater-than)
      - [Greater than or equal](#greater-than-or-equal)
      - [Less than](#less-than)
      - [Less than or equal](#less-than-or-equal)
      - [Like](#like)
      - [In](#in)
      - [Between](#between)
      - [Is null](#is-null)
      - [Grouping](#grouping)
      - [Not](#not)
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
```sql
SELECT id, name, email FROM users LIMIT 10;
```
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

// SelectUsers selects users.
func SelectUsers(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
		clauses.Selectable{Expr: expressions.Column{Name: "name"}},
		clauses.Selectable{Expr: expressions.Column{Name: "email"}},
	).
		FromTable("addresses", nil).
		Limit(args.NewArgument(10)).
		Write(sb)
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
```sql
SELECT country, COUNT(id) AS address_count
FROM addresses
GROUP BY country
ORDER BY address_count DESC;
```
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
	err := gosqle.NewSelect(
		clauses.Selectable{
			Expr: &expressions.Column{Name: "country"},
		},
		clauses.Selectable{
			Expr: expressions.NewCount(&expressions.Column{Name: "id"}),
			As:   "address_count",
		},
	).FromTable("addresses", nil).GroupBy(groupby.ColumnGrouping{
		&expressions.Column{Name: "country"},
	}).OrderBy(orderby.Sort{
		Column:    &expressions.Column{Name: "address_count"},
		Direction: orderby.DESC,
	}).Write(sb)
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
```sql
SELECT name
FROM users
WHERE id IN (
  SELECT user_id
  FROM addresses
  WHERE city = 'Amsterdam'
);
```
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
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "name"}},
	).From(
		from.NewFrom("users", ""),
	).Where(
		predicates.In{
			Col: expressions.Column{Name: "id"},
			Expr: gosqle.NewSelect(
				clauses.Selectable{Expr: expressions.Column{Name: "user_id"}},
			).From(
				from.NewFrom("addresses", ""),
			).Where(predicates.EQ{
				Col:  expressions.Column{Name: "city"},
				Expr: args.NewArgument("Amsterdam"),
			}).Statement, // <- This is the subquery, so without semicolon.
		},
	).Write(sb)

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
```sql
INSERT INTO users (name, email) VALUES ($1, $2)
```
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
	err := gosqle.NewInsert("users", "name", "email").Values(
		args.NewArgument("John"),
		args.NewArgument(fmt.Sprintf("john%d@%s", time.Now().Unix(), "example.com")),
	).Write(sb)

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
```sql
DELETE FROM users WHERE id = $1
```
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
	err := gosqle.NewDelete("addresses").Where(predicates.EQ{
		Col:  expressions.Column{Name: "user_id"},
		Expr: args.NewArgument(111),
	}).Write(sb)
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
```sql
UPDATE users SET name = $1 WHERE id = $2
```
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
	err := gosqle.NewUpdate("users").Set(set.Change{
		Col:  "name",
		Expr: args.NewArgument(fmt.Sprintf("new name %d", time.Now().Unix())),
	}).Where(predicates.EQ{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(1),
	}).Write(sb)
	if err != nil {
		return "", err
	}
	if _, err = db.Exec(sb.String(), args.Args...); err != nil {
		return "", err
	}
	return sb.String(), nil
}

```

### Where conditions
#### equal
```sql
SELECT id FROM users WHERE name = $1;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereEQ selects users where name is equal to 'John'.
func WhereEQ(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.EQ{
		Col:  expressions.Column{Name: "name"},
		Expr: args.NewArgument("John"),
	}).Write(sb)

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

```

#### Not equal
```sql
SELECT id FROM users WHERE name != $1;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereNE selects users where name is not equal to 'John'.
func WhereNE(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.NE{
		Col:  expressions.Column{Name: "name"},
		Expr: args.NewArgument("John"),
	}).Write(sb)

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

```

#### Greater than
```sql
SELECT id FROM users WHERE id > $1;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereGT selects users where id is greater than 10
func WhereGT(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.GT{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(10),
	}).Write(sb)

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

```

#### Greater than or equal
```sql
SELECT id FROM users WHERE id >= $1;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereGTE selects users where id is greater than or equal to 10
func WhereGTE(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.GTE{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(10),
	}).Write(sb)

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

```

#### Less than
```sql
SELECT id FROM users WHERE id &lt; $1;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLT selects users where id is less than 10
func WhereLT(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.LT{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(10),
	}).Write(sb)

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

```

#### Less than or equal
```sql
SELECT id FROM users WHERE id &lt;= $1;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLTE selects users where id is less than or equal to 10
func WhereLTE(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.LT{
		Col:  expressions.Column{Name: "id"},
		Expr: args.NewArgument(10),
	}).Write(sb)

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

```

#### Like
```sql
SELECT id FROM users WHERE name LIKE $1;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLike selects users where name is like anna%
func WhereLike(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.Like{
		Col:  expressions.Column{Name: "name"},
		Expr: args.NewArgument("anna%"),
	}).Write(sb)

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

```

#### In
```sql
SELECT id FROM users WHERE name IN ($1, $2, $3);
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereIN selects users where name is in 'John', 'Jane' or 'Joe'.
func WhereIN(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.In{
		Col: expressions.Column{Name: "id"},
		Expr: expressions.List{
			args.NewArgument("John"),
			args.NewArgument("Jane"),
			args.NewArgument("Joe"),
		},
	}).Write(sb)

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

```

#### Between 
```sql
SELECT id FROM users WHERE id BETWEEN $1 AND $2;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereBetween selects users where id is between 10 and 20
func WhereBetween(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).
		Where(predicates.Between{
			Col:  expressions.Column{Name: "id"},
			Low:  args.NewArgument(10),
			High: args.NewArgument(20),
		}).Write(sb)

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

```

#### Is null
```sql
SELECT id FROM addresses WHERE phone IS NULL;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereIsNull selects addresses where phone is null
func WhereIsNull(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.IsNull{
		Col: expressions.Column{Name: "phone"},
	}).Write(sb)

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

```

#### Grouping
```sql
SELECT id FROM users WHERE (id BETWEEN $1 AND $2 OR id BETWEEN $3 AND $4) OR name = $5;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereWrap selects users where id is between 10 and 20 or 30 and 40 or name is john
func WhereWrap(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(
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
	).Write(sb)

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

```

#### Not
```sql
SELECT id FROM users WHERE NOT name = $1;
```
```go
package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereNOT selects users where name is not John.
func WhereNOT(db *sql.DB) ([]User, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.Column{Name: "id"}},
	).FromTable("users", nil).Where(predicates.Not{
		Predicate: predicates.EQ{
			Col:  expressions.Column{Name: "name"},
			Expr: args.NewArgument("John"),
		},
	}).Write(sb)

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

```

## Syntax used
![image](provision/images/SQL_syntax.svg)