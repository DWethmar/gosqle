# gosqle<!-- Don't edit README.md, but edit README.tmpl.md as that one is used to generate the README.md -->
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
{{insertGoFile "examples/select.go" }}
```

#### Generate select query using group by and aggregate functions:
```sql
SELECT country, COUNT(id) AS address_count
FROM addresses
GROUP BY country
ORDER BY address_count DESC;
```
```go
{{insertGoFile "examples/select-aggregate.go" }}
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
{{insertGoFile "examples/subquery.go" }}
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
{{insertGoFile "examples/insert.go" }}
```

### Delete
#### Generate a delete query:
```sql
DELETE FROM users WHERE id = $1
```
```go
{{insertGoFile "examples/delete.go" }}
```

### Update
#### Generate an update query:
```sql
UPDATE users SET name = $1 WHERE id = $2
```
```go
{{insertGoFile "examples/update.go" }}
```

### Where conditions
#### equal
```sql
SELECT id FROM users WHERE name = $1;
```
```go
{{insertGoFile "examples/where-eq.go" }}
```

#### Not equal
```sql
SELECT id FROM users WHERE name != $1;
```
```go
{{insertGoFile "examples/where-ne.go" }}
```

#### Greater than
```sql
SELECT id FROM users WHERE id > $1;
```
```go
{{insertGoFile "examples/where-gt.go" }}
```

#### Greater than or equal
```sql
SELECT id FROM users WHERE id >= $1;
```
```go
{{insertGoFile "examples/where-gte.go" }}
```

#### Less than
```sql
SELECT id FROM users WHERE id < $1;
```
```go
{{insertGoFile "examples/where-lt.go" }}
```

#### Less than or equal
```sql
SELECT id FROM users WHERE id <= $1;
```
```go
{{insertGoFile "examples/where-lte.go" }}
```

#### Like
```sql
SELECT id FROM users WHERE name LIKE $1;
```
```go
{{insertGoFile "examples/where-like.go" }}
```

#### In
```sql
SELECT id FROM users WHERE name IN ($1, $2, $3);
```
```go
{{insertGoFile "examples/where-in.go" }}
```

#### Between 
```sql
SELECT id FROM users WHERE id BETWEEN $1 AND $2;
```
```go
{{insertGoFile "examples/where-between.go" }}
```

#### Is null
```sql
SELECT id FROM addresses WHERE phone IS NULL;
```
```go
{{insertGoFile "examples/where-is-null.go" }}
```

#### Grouping
```sql
SELECT id FROM users WHERE (id BETWEEN $1 AND $2 OR id BETWEEN $3 AND $4) OR name = $5;
```
```go
{{insertGoFile "examples/where-wrap.go" }}
```

#### Not
```sql
SELECT id FROM users WHERE NOT name = $1;
```
```go
{{insertGoFile "examples/where-not.go" }}
```

## Syntax used
![image](provision/images/SQL_syntax.svg)