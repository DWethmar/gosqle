package examples

import (
	"testing"
)

func TestPeopleOfAmsterdam(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	if db == nil {
		t.Error("db is nil")
		return
	}

	t.Run("select people of amsterdam should render correct sql", func(t *testing.T) {
		const want = "SELECT id, name, email FROM users WHERE id IN (SELECT user_id FROM addresses WHERE city = $1);"

		args, query, err := PeopleOfCity("Amsterdam")
		if err != nil {
			t.Errorf("error selecting people of amsterdam: %v", err)
		}
		if query != want {
			t.Errorf("query is not correct, got: %q, want: %q", query, want)
		}
		if len(args) != 1 {
			t.Errorf("args length is not correct, got: %q, want: %q", len(args), 1)
		}
		if args[0] != "Amsterdam" {
			t.Errorf("first arg is not correct, got: %q, want: %q", args[0], "Amsterdam")
		}
	})

	t.Run("select people of amsterdam should execute", func(t *testing.T) {
		args, query, err := PeopleOfCity("Amsterdam")
		if err != nil {
			t.Errorf("error selecting people of amsterdam: %v", err)
		}

		rows, err := db.Query(query, args...)
		if err != nil {
			t.Errorf("error selecting people of amsterdam: %v", err)
		}

		defer rows.Close()

		for rows.Next() {
			var id int64
			var name string
			var email string
			if err := rows.Scan(&id, &name, &email); err != nil {
				t.Errorf("error selecting people of amsterdam: %v", err)
			}

			t.Logf("id: %d, name: %q, email: %q", id, name, email)
		}
	})
}
