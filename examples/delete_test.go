package examples

import (
	"database/sql"
	"testing"
)

func TestExampleDeleteAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	if db == nil {
		t.Error("db is nil")
		return
	}

	t.Run("delete address should render correct sql", func(t *testing.T) {
		const id = int64(1)
		const want = "DELETE FROM addresses WHERE addresses.id = $1;"

		args, query, err := DeleteAddress(id)
		if err != nil {
			t.Errorf("error deleting address: %v", err)
		}
		if query != want {
			t.Errorf("got: %q, want: %q", query, want)
		}
		if len(args) != 1 {
			t.Errorf("got: %q, want: %q", len(args), id)
		}

		if args[0] != id {
			t.Errorf("got: %q, want: %q", args[0], id)
		}
	})

	t.Run("delete address should execute", func(t *testing.T) {
		userId, err := InsertAddress(db, &Address{
			RecipientName: "John Doe",
			AddressLine1:  "Street 1",
			AddressLine2: sql.NullString{
				String: "Street 2",
				Valid:  true,
			},
			AddressLine3: sql.NullString{
				String: "Street 3",
				Valid:  true,
			},
			City: "City",
			StateProvinceRegion: sql.NullString{
				String: "State",
				Valid:  true,
			},
			PostalCode: sql.NullString{},
			Country:    "Country",
			Phone:      sql.NullString{},
		})

		if err != nil {
			t.Errorf("error inserting address: %v", err)
		}

		args, query, err := DeleteAddress(userId)
		if err != nil {
			t.Errorf("error deleting address: %v", err)
		}

		r, err := db.Exec(query, args...)
		if err != nil {
			t.Errorf("error deleting address: %v", err)
		}

		affected, err := r.RowsAffected()
		if err != nil {
			t.Errorf("error deleting address: %v", err)
		}

		if affected != 1 {
			t.Error("expected 1 row to be affected")
		}
	})
}
