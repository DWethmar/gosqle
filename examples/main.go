package main

import (
	"database/sql"
	"flag"
	"fmt"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

var examples = map[string]func(*sql.DB) error{
	"select": func(db *sql.DB) error {
		args, query, err := SelectUsers()
		if err != nil {
			return fmt.Errorf("error selecting users: %v", err)
		}
		fmt.Printf("Query: %q\n", query)
		result, err := QueryUsers(db, args, query)
		if err != nil {
			return fmt.Errorf("error executing query: %v", err)
		}
		PrintUsers(result)
		return nil
	},
	"select-aggregate": func(db *sql.DB) error {
		query, err := SelectAmountOfAddressesPerCountry()
		if err != nil {
			fmt.Printf("error selecting users: %v\n", err)
		}
		fmt.Printf("Query: %q\n", query)
		rows, err := db.Query(query)
		if err != nil {
			return fmt.Errorf("error executing query: %v", err)
		}
		defer rows.Close()
		for rows.Next() {
			var u AmountOfAddressesPerCountry
			if err := rows.Scan(&u.Country, &u.AddressCount); err != nil {
				return fmt.Errorf("error scanning row: %v", err)
			}
			fmt.Printf("Country: %s, AddressCount: %d\n", u.Country, u.AddressCount)
		}
		return nil
	},
	"subquery": func(d *sql.DB) error {
		args, query, err := PeopleOfAmsterdam()
		if err != nil {
			fmt.Printf("error selecting users: %v\n", err)
		}
		fmt.Printf("Query: %q\n", query)
		result, err := QueryUsers(d, args, query)
		if err != nil {
			fmt.Printf("error executing query: %v\n", err)
		}
		PrintUsers(result)
		return nil
	},
	"insert": func(d *sql.DB) error {
		args, query, err := InsertUser()
		if err != nil {
			fmt.Printf("error inserting user: %v\n", err)
		}
		fmt.Printf("Query: %q\n", query)
		result, err := Exec(d, args, query)
		if err != nil {
			fmt.Printf("error executing query: %v\n", err)
		}
		fmt.Printf("Result: %v\n", result)
		return nil
	},
	"update": func(d *sql.DB) error {
		query, err := UpdateUser()
		if err != nil {
			fmt.Printf("error updating user: %v\n", err)
		}
		fmt.Printf("Query: %q\n", query)
		result, err := Exec(d, []interface{}{}, query)
		if err != nil {
			fmt.Printf("error executing query: %v\n", err)
		}
		fmt.Printf("Result: %v\n", result)
		return nil
	},
	"delete": func(d *sql.DB) error {
		args, query, err := DeleteAddress()
		if err != nil {
			fmt.Printf("error deleting address: %v\n", err)
		}
		fmt.Printf("Query: %q\n", query)
		result, err := Exec(d, args, query)
		if err != nil {
			fmt.Printf("error executing query: %v\n", err)
		}
		fmt.Printf("Result: %v\n", result)
		return nil
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
		if err := f(db); err != nil {
			fmt.Printf("Error running example %q: %v\n", example, err)
		}
	} else {
		exampleNames := make([]string, 0, len(examples))
		for k := range examples {
			exampleNames = append(exampleNames, k)
		}
		sort.Strings(exampleNames)
		fmt.Printf("Example %s not found, available examples: %v\n", example, strings.Join(exampleNames, ", "))
	}
}

func QueryUsers(db *sql.DB, args []interface{}, query string) ([]User, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		result = append(result, u)
	}

	return result, nil
}

func Exec(db *sql.DB, args []interface{}, query string) ([]interface{}, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	var r []interface{}
	if id, err := result.LastInsertId(); err == nil {
		r = append(r, id)
	}
	if rows, err := result.RowsAffected(); err == nil {
		r = append(r, rows)
	}

	return r, nil
}

func PrintUsers(l []User) {
	for i, r := range l {
		fmt.Printf("#%d %+v\n", i, r)
	}
}
