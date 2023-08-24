package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Define the flag. The default value for the database URL is an empty string.
	var dbURL string
	flag.StringVar(&dbURL, "dburl", "", "The URL of the database to connect to")

	// Parse the flags
	flag.Parse()

	if dbURL == "" {
		fmt.Println("Please provide a database URL using the -dburl flag.")
		return
	}

	fmt.Println("Database URL:", dbURL)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	// truncate the tables
	if err := truncate(db); err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	defer func() {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			if errr := tx.Rollback(); errr != nil {
				log.Fatal(errr)
			}
		} else {
			if errr := tx.Commit(); errr != nil {
				log.Fatal(errr)
			}
		}
	}()

	if err = seedUsers(tx); err != nil {
		fmt.Printf("Error seeding users: %v\n", err)
		return
	}

	if err = seedAddressTypes(tx); err != nil {
		fmt.Printf("Error seeding address types: %v\n", err)
		return
	}

	if err = seedAddresses(tx); err != nil {
		fmt.Printf("Error seeding addresses: %v\n", err)
		return
	}
}

func truncate(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// truncate the address table
	_, err = tx.Exec("TRUNCATE TABLE addresses CASCADE")
	if err != nil {
		return fmt.Errorf("error truncating addresses table: %v", err)
	}

	// truncate the users table
	_, err = tx.Exec("TRUNCATE TABLE users CASCADE")
	if err != nil {
		return fmt.Errorf("error truncating users table: %v", err)
	}

	// truncate the address_type
	_, err = tx.Exec("TRUNCATE TABLE address_types CASCADE")
	if err != nil {
		return fmt.Errorf("error truncating address_types table: %v", err)
	}

	return tx.Commit()
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
