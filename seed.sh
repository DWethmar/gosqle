#! /bin/bash

go run scripts/seed/*.go -dburl="postgres://postgres:postgres@localhost:5439/customers?sslmode=disable"
