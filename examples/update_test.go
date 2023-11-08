package examples

import (
	"testing"
)

func TestUpdateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	if db == nil {
		t.Error("db is nil")
		return
	}

	t.Run("update user should render correct sql", func(t *testing.T) {
		const want = "UPDATE users SET name = $1 WHERE id = $2;"

		newName := "Dennis"
		args, query, err := UpdateUser(1, newName)

		if err != nil {
			t.Errorf("error updating user: %v", err)
		}
		if query != want {
			t.Errorf("query is not correct, got: %q, want: %q", query, want)
		}
		if len(args) != 2 {
			t.Errorf("args length is not correct, got: %q, want: %q", len(args), 2)
		}
		if args[0] != newName {
			t.Errorf("first arg is not correct, got: %q, want: %q", args[0], newName)
		}
		if args[1] != int64(1) {
			t.Errorf("second arg is not correct, got: %q, want: %q", args[1], 1)
		}
	})

	t.Run("update user should execute", func(t *testing.T) {
		newName := "Dennis"
		args, query, err := UpdateUser(1, newName)
		if err != nil {
			t.Errorf("error updating user: %v", err)
		}
		result, err := db.Exec(query, args...)
		if err != nil {
			t.Errorf("error updating user: %v", err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			t.Errorf("error updating user: %v", err)
		}
		if rowsAffected != int64(1) {
			t.Errorf("rows affected is not correct, got: %d, want: %d", rowsAffected, int64(1))
		}
	})
}
