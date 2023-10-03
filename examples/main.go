package main

import (
	"database/sql"
	"flag"
	"fmt"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

var examples = map[string]func(*sql.DB){
	"select": func(d *sql.DB) {
		users, query, err := SelectUsers(d)
		if err != nil {
			fmt.Printf("error selecting users: %v\n", err)
		}
		fmt.Printf("Query: %q\n", query)
		for i, u := range users {
			fmt.Printf("#%d %+v\n", i, u)
		}
	},
	"select-aggregate": func(d *sql.DB) {
		amounts, query, err := SelectAmountOfAddressesPerCountry(d)
		if err != nil {
			fmt.Printf("error selecting amounts: %v\n", err)
		}
		fmt.Printf("Query: %q\n", query)
		for i, a := range amounts {
			fmt.Printf("#%d %+v\n", i, a)
		}
	},
	"subquery": func(d *sql.DB) {
		users, query, err := PeopleOfAmsterdam(d)
		if err != nil {
			fmt.Printf("error selecting users: %v\n", err)
		}
		fmt.Printf("Query: %q\n", query)
		for i, u := range users {
			fmt.Printf("#%d %+v\n", i, u)
		}
	},
	"insert": func(d *sql.DB) {
		if q, err := InsertUser(d); err == nil {
			fmt.Printf("Query: %q\n", q)
		} else {
			fmt.Printf("error inserting user: %v\n", err)
		}
	},
	"update": func(d *sql.DB) {
		if q, err := UpdateUser(d); err == nil {
			fmt.Printf("Query: %q\n", q)
		} else {
			fmt.Printf("error updating user: %v\n", err)
		}
	},
	"delete": func(d *sql.DB) {
		if q, err := DeleteAddress(d); err == nil {
			fmt.Printf("Query: %q\n", q)
		} else {
			fmt.Printf("error deleting user: %v\n", err)
		}
	},
}

func main() {
	var dbURL string
	var example string

	flag.StringVar(&dbURL, "dburl", "", "The URL of the database to connect to")
	flag.StringVar(&example, "example", "", "The example to run")
	flag.Parse()

	if dbURL == "" {
		fmt.Println("Please provide a database URL using the -dburl flag.")
		return
	}

	if example == "" {
		fmt.Println("Please provide an example to run using the -example flag.")
		return
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	if f, ok := examples[example]; ok {
		fmt.Printf("Running example %q:\n", example)
		f(db)
	} else {
		exampleNames := make([]string, 0, len(examples))
		for k := range examples {
			exampleNames = append(exampleNames, k)
		}
		sort.Strings(exampleNames)
		fmt.Printf("Example %s not found, available examples: %v\n", example, strings.Join(exampleNames, ", "))
	}
}
