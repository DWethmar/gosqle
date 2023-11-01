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
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

// SelectUsers selects users.
func SelectUsers() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
		alias.New(expressions.Column{Name: "name"}),
		alias.New(expressions.Column{Name: "email"}),
	).
		FromTable("users", nil).
		Limit(args.NewArgument(10)).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
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
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/clauses/groupby"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/expressions"
)

// SelectAmountOfAddressesPerCountry select amount of addresses per country
func SelectAmountOfAddressesPerCountry() (string, error) {
	sb := new(strings.Builder)
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "country"}),
		alias.New(expressions.Column{Name: "address_count"}).SetAs("address_count"),
	).FromTable("addresses", nil).
		GroupBy(groupby.ColumnGrouping{
			&expressions.Column{Name: "country"},
		}).
		OrderBy(orderby.Sort{
			Column:    &expressions.Column{Name: "address_count"},
			Direction: orderby.DESC,
		}).Write(sb)

	if err != nil {
		return "", err
	}

	return sb.String(), nil
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
	"fmt"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// SelectUsers selects users.
func PeopleOfAmsterdam() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "name"}),
	).FromTable("users", nil).
		Where(
			logic.And(
				predicates.IN(
					expressions.Column{Name: "city"},
					gosqle.NewSelect(
						alias.New(expressions.Column{Name: "id"}),
					).Where(
						logic.And(predicates.EQ(expressions.Column{Name: "id", From: "users"}, args.NewArgument(1))),
					).Statement, // <-- This is the sub-query without semicolon
				),
			),
		).Write(sb)

	if err != nil {
		return nil, "", fmt.Errorf("error writing query: %v", err)
	}

	return args.Values, sb.String(), nil
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

	return args.Values, sb.String(), nil
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
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// Delete deletes a user.
func DeleteAddress() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewDelete("addresses").
		Where(
			logic.And(predicates.EQ(expressions.Column{Name: "id", From: "addresses"}, args.NewArgument(1))),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
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
	"fmt"
	"strings"
	"time"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses/set"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// UpdateUser updates a user.
func UpdateUser() (string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewUpdate("users").Set(set.Change{
		Col:  "name",
		Expr: args.NewArgument(fmt.Sprintf("new name %d", time.Now().Unix())),
	}).Where(
		logic.And(predicates.EQ(
			expressions.Column{Name: "id"},
			args.NewArgument(1),
		)),
	).Write(sb)
	if err != nil {
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
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereEQ selects users where name is equal to 'John'.
func WhereEQ(username string) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.EQ(
				expressions.Column{Name: "name"},
				args.NewArgument(username),
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### Not equal
```sql
SELECT id FROM users WHERE name != $1;
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereNE selects users where name is not equal to 'John'.
func WhereNE() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.NE(
				expressions.Column{Name: "name"},
				args.NewArgument("John"),
			)),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### Greater than
```sql
SELECT id FROM users WHERE id > $1;
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereGT selects users where id is greater than 10
func WhereGT(id int) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.GT(
				expressions.Column{Name: "id"},
				args.NewArgument(id),
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### Greater than or equal
```sql
SELECT id FROM users WHERE id >= $1;
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereGTE selects users where id is greater than or equal to 10
func WhereGTE(id int) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.GTE(
				expressions.Column{Name: "id"},
				args.NewArgument(id),
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### Less than
```sql
SELECT id FROM users WHERE id &lt; $1;
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLT selects users where id is less than 10
func WhereLT() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.LT(
				expressions.Column{Name: "id"},
				args.NewArgument(10),
			)),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### Less than or equal
```sql
SELECT id FROM users WHERE id &lt;= $1;
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLTE selects users where id is less than or equal to 10
func WhereLTE() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.LTE(
				expressions.Column{Name: "id"},
				args.NewArgument(10),
			)),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### Like
```sql
SELECT id FROM users WHERE name LIKE $1;
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereLike selects users where name is like anna%
func WhereLike() ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(logic.And(predicates.Like(
			expressions.Column{Name: "name"},
			args.NewArgument("anna%"),
		))).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### In
```sql
SELECT id FROM users WHERE name IN ($1, $2, $3);
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereIN selects users where name is in names.
func WhereIN(names []string) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()

	list := []expressions.Expression{}
	for _, name := range names {
		list = append(list, args.NewArgument(name))
	}

	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.IN(
				expressions.Column{Name: "name"},
				list...,
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### Between 
```sql
SELECT id FROM users WHERE id BETWEEN $1 AND $2;
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereBetween selects users where id is between 10 and 20
func WhereBetween(low, high int) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.Between(
				expressions.Column{Name: "id"},
				args.NewArgument(low),
				args.NewArgument(high),
			)),
		).Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

#### Is null
```sql
SELECT id FROM addresses WHERE phone IS NULL;
```
```go
package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereIsNull selects addresses where phone is null
func WhereIsNull() (string, error) {
	sb := new(strings.Builder)
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.IsNull(expressions.Column{Name: "phone"})),
		).Write(sb)

	if err != nil {
		return "", err
	}

	return sb.String(), nil
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
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereWrap selects users where id is between 10 and 20 or 30 and 40 or name is john
func WhereWrap(db *sql.DB) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.Between(
				expressions.Column{Name: "id"},
				args.NewArgument(10),
				args.NewArgument(20),
			)),
			logic.Or(
				logic.Group([]logic.Logic{
					logic.And(predicates.Between(
						expressions.Column{Name: "id"},
						args.NewArgument(30),
						args.NewArgument(40),
					)),
					logic.Or(predicates.EQ(
						expressions.Column{Name: "name"},
						args.NewArgument("John"),
					)),
				}),
			),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
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
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereNOT selects users where name is not John.
func WhereNOT(db *sql.DB) ([]interface{}, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).FromTable("users", nil).
		Where(
			logic.And(predicates.Not(predicates.EQ(
				expressions.Column{Name: "name"},
				args.NewArgument("John"),
			))),
		).
		Write(sb)

	if err != nil {
		return nil, "", err
	}

	return args.Values, sb.String(), nil
}

```

## Syntax used
![image](provision/images/SQL_syntax.svg)