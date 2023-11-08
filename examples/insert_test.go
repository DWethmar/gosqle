package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestExampleInsertUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	if db == nil {
		t.Error("db is nil")
		return
	}

	t.Run("insert a new user", func(t *testing.T) {
		const name = "John"
		var email = fmt.Sprintf("%s%d@test.test", name, time.Now().Unix())

		const want = "INSERT INTO users (name, email) VALUES ($1, $2);"

		args, query, err := InsertUser(name, email)
		if err != nil {
			t.Errorf("error inserting user: %v", err)
		}

		if query != want {
			t.Errorf("got: %v, want: %v", query, want)
		}

		if len(args) != 2 {
			t.Errorf("got: %v, want: %v", len(args), 2)
		}

		if args[0] != name {
			t.Errorf("got: %v, want: %v", args[0], name)
		}

		if args[1] != email {
			t.Errorf("got: %v, want: %v", args[1], email)
		}
	})

	t.Run("insert a new user should execute", func(t *testing.T) {
		const name = "John"
		var email = fmt.Sprintf("%s%d@test.test", name, time.Now().Unix())

		args, query, err := InsertUser(name, email)
		if err != nil {
			t.Errorf("error inserting user: %v", err)
		}

		if _, err := db.Exec(query, args...); err != nil {
			t.Errorf("error inserting user: %v", err)
		}
	})
}
