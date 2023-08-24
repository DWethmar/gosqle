#! /bin/bash
dburl="postgres://postgres:postgres@localhost:5439/customers?sslmode=disable"

go run examples/*.go -dburl=$dburl -example="select"
go run examples/*.go -dburl=$dburl -example="insert"
