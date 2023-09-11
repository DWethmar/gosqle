package where

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestWhere(t *testing.T) {
	type args struct {
		sb          *strings.Builder
		predicates  []predicates.Predicate
		paramOffset int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "empty filter",
			args: args{
				sb:          &strings.Builder{},
				predicates:  []predicates.Predicate{},
				paramOffset: 0,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "where clause with one condition",
			args: args{
				sb: &strings.Builder{},
				predicates: []predicates.Predicate{
					predicates.EQ{
						Col:  expressions.Column{Name: "id"},
						Expr: postgres.NewArgument(1, 1),
					},
				},
				paramOffset: 0,
			},
			want:    `WHERE id = $1`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteWhere(tt.args.sb, tt.args.predicates); (err != nil) != tt.wantErr {
				t.Errorf("Where() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb := tt.args.sb; sb != nil {
				if str := sb.String(); str != tt.want {
					t.Errorf("Where() got = %q, want %q", str, tt.want)
				}
			}
		})
	}

	t.Run("should return io.writer error", func(t *testing.T) {
		writer := mock.StringWriterFn(func(s string) (n int, err error) {
			return 0, errors.New("error")
		})

		if err := WriteWhere(writer, []predicates.Predicate{predicates.EQ{}}); err == nil {
			t.Error("expected error")
		}
	})
}
