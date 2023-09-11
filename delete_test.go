package gosqle

import (
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mysql"
	"github.com/dwethmar/gosqle/predicates"
)

func TestDelete_WriteTo(t *testing.T) {
	tests := []struct {
		name    string
		delete  *Delete
		want    string
		wantErr bool
	}{
		{
			name: "delete from",
			delete: NewDelete("users").Where(
				predicates.EQ{
					Col:  expressions.Column{Name: "id"},
					Expr: mysql.NewArgument(1),
				},
			),
			want: "DELETE FROM users WHERE id = ?;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := new(strings.Builder)
			err := tt.delete.WriteTo(sb)

			if (err != nil) != tt.wantErr {
				t.Errorf("Delete.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if query := sb.String(); query != tt.want {
				t.Errorf("Delete.WriteTo() query = %q, wantQuery %q", query, tt.want)
			}
		})
	}
}

func TestNewDelete(t *testing.T) {
	type args struct {
		table string
	}
	tests := []struct {
		name string
		args args
		want *Update
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDelete(tt.args.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
