#! /bin/bash
dburl="postgres://postgres:postgres@localhost:5439/customers?sslmode=disable"

go run scripts/seed/*.go -dburl=$dburl
