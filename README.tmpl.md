# gosqle<!-- Don't edit README.md, but edit README.tmpl.md as that one is used to generate the README.md -->
gosqle is a golang package that can generate sql queries. 

Table of Contents:
- [gosqle](#gosqle)
  - [Examples](#examples)
    - [Select statement](#select-statement)
    - [Insert statement](#insert-statement)
    - [Delete statement](#delete-statement)
    - [Update statement](#update-statement)
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
Generate a select query:
```go
{{insertGoFile "examples/select.go" }}
```

### Insert statement
Generate a insert query:
```go
{{insertGoFile "examples/insert.go" }}
```

### Delete statement
Generate a delete query:
```go
{{insertGoFile "examples/delete.go" }}
```

### Update statement
Generate a update query:
```go
{{insertGoFile "examples/update.go" }}
```

## Syntax used

![image](provision/images/SQL_syntax.svg)