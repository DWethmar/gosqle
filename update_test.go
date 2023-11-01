package gosqle

import (
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses/set"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestUpdate_Write(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		update  *Update
		want    string
		wantErr bool
	}{
		{
			name: "update",
			update: NewUpdate("users").Set(
				set.Change{Col: "name", Expr: postgres.NewArgument("John", 1)},
				set.Change{Col: "age", Expr: postgres.NewArgument(25, 2)},
			).Where(
				logic.And(predicates.EQ(
					expressions.Column{Name: "id"},
					postgres.NewArgument(1, 3),
				)),
			),
			want: "UPDATE users SET name = $1, age = $2 WHERE id = $3;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := new(strings.Builder)
			err := tt.update.Write(sb)

			if (err != nil) != tt.wantErr {
				t.Errorf("Update.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if query := sb.String(); query != tt.want {
				t.Errorf("Update.Write() query = %q, wantQuery %q", query, tt.want)
			}
		})
	}
}
