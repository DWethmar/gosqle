# gosqle<!-- Don't edit README.md, but edit README.tmpl.md as that one is used to generate the README.md -->
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
{{insertGoFile "examples/select.go" }}
```

## Insert statement
Generate a insert query:
```go
{{insertGoFile "examples/insert.go" }}
```

## Delete statement
Generate a delete query:
```go
{{insertGoFile "examples/delete.go" }}
```

## Syntax used

![image](provision/images/SQL_syntax.svg)