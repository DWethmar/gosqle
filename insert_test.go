package gosqle

import (
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/mysql"
	"github.com/dwethmar/gosqle/postgres"
)

func TestInsert_WriteTo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		insert    *Insert
		wantQuery string
		wantErr   bool
	}{
		{
			name: "select columns mysql",
			insert: NewInsert("users", "id", "username").Values(
				mysql.NewArgument(1),
				mysql.NewArgument("test"),
			),
			wantQuery: "INSERT INTO users (id, username) VALUES (?, ?);",
			wantErr:   false,
		},
		{
			name: "select columns postgres",
			insert: NewInsert("users", "id", "username").Values(
				postgres.NewArgument(1, 1),
				postgres.NewArgument("test", 2),
			),
			wantQuery: "INSERT INTO users (id, username) VALUES ($1, $2);",
			wantErr:   false,
		},
		{
			name:      "error on no table",
			insert:    NewInsert("", "email", "username"),
			wantQuery: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := new(strings.Builder)
			err := tt.insert.WriteTo(sb)

			if (err != nil) != tt.wantErr {
				t.Errorf("Insert.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if query := sb.String(); query != tt.wantQuery {
				t.Errorf("Insert.WriteTo() query = %q, wantQuery %q", query, tt.wantQuery)
			}
		})
	}
}

func BenchmarkInsert_WriteTo(b *testing.B) {
	insert := NewInsert("users", "id", "username").Values(
		mysql.NewArgument(1),
		mysql.NewArgument("test"),
	)

	sb := new(strings.Builder)
	for i := 0; i < b.N; i++ {
		if err := insert.WriteTo(sb); err != nil {
			b.Fatal(err)
		}
	}
}
