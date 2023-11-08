package examples

import (
	"testing"
)

func TestExampleSelectAmountOfAddressesPerCountry(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	if db == nil {
		t.Error("db is nil")
		return
	}

	t.Run("select amount of addresses per country should render correct sql", func(t *testing.T) {
		const want = "SELECT country, COUNT(id) AS address_count FROM addresses GROUP BY country ORDER BY address_count DESC;"

		query, err := SelectAmountOfAddressesPerCountry()
		if err != nil {
			t.Errorf("error selecting amount of addresses per country: %v", err)
		}
		if query != want {
			t.Errorf("got: %q, want: %q", query, want)
		}
	})

	t.Run("select amount of addresses per country should execute", func(t *testing.T) {
		query, err := SelectAmountOfAddressesPerCountry()
		if err != nil {
			t.Errorf("error selecting amount of addresses per country: %v", err)
		}

		rows, err := db.Query(query)
		if err != nil {
			t.Errorf("error selecting amount of addresses per country: %v", err)
		}

		defer rows.Close()
		for rows.Next() {
			var country string
			var addressCount int
			if err := rows.Scan(&country, &addressCount); err != nil {
				t.Errorf("error selecting amount of addresses per country: %v", err)
			}

			t.Logf("country: %s count: %d", country, addressCount)
		}
	})
}
