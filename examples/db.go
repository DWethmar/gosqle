package examples

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var dbURL string
	var err error

	dbURL = os.Getenv("TEST_DB_URL")

	if dbURL == "" {
		fmt.Println("TEST_DB_URL not set")
		return
	}

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
}
