# gosqle<!-- Don't edit README.md, but edit README.tmpl.md as that one is used to generate the README.md -->
gosqle is a golang package that can generate sql queries. 

Table of Contents:
- [gosqle](#gosqle)
  - [Examples](#examples)
    - [Select statement](#select-statement)
      - [Generate a select query:](#generate-a-select-query)
      - [Generate select query using group by and aggregate functions:](#generate-select-query-using-group-by-and-aggregate-functions)
    - [Insert statement](#insert-statement)
      - [Generate an insert query:](#generate-an-insert-query)
    - [Delete statement](#delete-statement)
      - [Generate a delete query:](#generate-a-delete-query)
    - [Update statement](#update-statement)
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

### Select statement
Create a select statement with the following syntax:
```go
gosqle.NewSelect(...columns)
```
#### Generate a select query:
```go
{{insertGoFile "examples/select.go" }}
```

#### Generate select query using group by and aggregate functions:
```go
{{insertGoFile "examples/select-aggregate.go" }}
```

### Insert statement
```go
gosqle.NewInsert(table, ...columns)
```
#### Generate an insert query:
```go
{{insertGoFile "examples/insert.go" }}
```

### Delete statement
#### Generate a delete query:
```go
{{insertGoFile "examples/delete.go" }}
```

### Update statement
#### Generate an update query:
```go
{{insertGoFile "examples/update.go" }}
```

## Syntax used

![image](provision/images/SQL_syntax.svg)