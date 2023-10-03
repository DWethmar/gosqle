#! /bin/bash
dburl="postgres://postgres:postgres@localhost:5439/customers?sslmode=disable"

go run examples/*.go -dburl=$dburl -example="select"
go run examples/*.go -dburl=$dburl -example="select-aggregate"
go run examples/*.go -dburl=$dburl -example="subquery"
go run examples/*.go -dburl=$dburl -example="insert"
go run examples/*.go -dburl=$dburl -example="update"
go run examples/*.go -dburl=$dburl -example="delete"