package gosqle

import (
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mysql"
	"github.com/dwethmar/gosqle/predicates"
)

func TestDelete_Write(t *testing.T) {
	tests := []struct {
		name    string
		delete  *Delete
		want    string
		wantErr bool
	}{
		{
			name: "delete from",
			delete: NewDelete("users").Where(predicates.EQ{
				Col:  expressions.Column{Name: "id", From: "users"},
				Expr: mysql.NewArgument(1),
			}),
			want: "DELETE FROM users WHERE users.id = ?;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := new(strings.Builder)
			err := tt.delete.Write(sb)

			if (err != nil) != tt.wantErr {
				t.Errorf("Delete.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if query := sb.String(); query != tt.want {
				t.Errorf("Delete.Write() query = %q, wantQuery %q", query, tt.want)
			}
		})
	}
}
