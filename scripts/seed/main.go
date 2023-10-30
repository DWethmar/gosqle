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

func NewDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}

func main() {
	// Define the flag. The default value for the database URL is an empty string.
	var dbURL string
	flag.StringVar(&dbURL, "dburl", "", "The URL of the database to connect to")
	flag.Parse()

	if dbURL == "" {
		fmt.Println("Please provide a database URL using the -dburl flag.")
		return
	}

	db, err := NewDB(dbURL)
	if err != nil {
		panic(err)
	}

	// truncate the tables
	if err = truncate(db); err != nil {
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
		return fmt.Errorf("error beginning transaction: %v", err)
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

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return lines, nil
}
