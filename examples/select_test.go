package examples

import (
	"testing"
)

func TestSelectUsers(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	if db == nil {
		t.Error("db is nil")
		return
	}

	t.Run("select users should render correct sql", func(t *testing.T) {
		const want = "SELECT id, name, email FROM users LIMIT $1;"

		args, query, err := SelectUsers(10)
		if err != nil {
			t.Errorf("error selecting users: %v", err)
		}
		if query != want {
			t.Errorf("query is not correct, got: %q, want: %q", query, want)
		}
		if len(args) != 1 {
			t.Errorf("args length is not correct, got: %q, want: %q", len(args), 1)
		}
		if args[0] != int64(10) {
			t.Errorf("first arg is not correct, got: %q, want: %q", args[0], 10)
		}
	})

	t.Run("select users should execute", func(t *testing.T) {
		args, query, err := SelectUsers(10)
		if err != nil {
			t.Errorf("error selecting users: %v", err)
		}

		rows, err := db.Query(query, args...)
		if err != nil {
			t.Errorf("error selecting users: %v", err)
		}

		defer rows.Close()

		for rows.Next() {
			var id int64
			var name string
			var email string
			if err := rows.Scan(&id, &name, &email); err != nil {
				t.Errorf("error selecting users: %v", err)
			}

			t.Logf("id: %d, name: %q, email: %q", id, name, email)
		}
	})
}
