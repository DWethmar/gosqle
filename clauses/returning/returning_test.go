package returning

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
)

func TestWrite(t *testing.T) {
	type args struct {
		sb      *strings.Builder
		columns []clauses.Selectable
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return error when no fields are supplied",
			args: args{
				sb:      new(strings.Builder),
				columns: []clauses.Selectable{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should render returning clause",
			args: args{
				sb: new(strings.Builder),
				columns: []clauses.Selectable{
					{
						Expr: expressions.Column{Name: "field_a", From: "table_a"},
					},
					{
						Expr: expressions.Column{Name: "field_b", From: "table_b"},
					},
				},
			},
			want:    "RETURNING table_a.field_a, table_b.field_b",
			wantErr: false,
		},
		{
			name: "should render returning clause with has alias",
			args: args{
				sb: new(strings.Builder),
				columns: []clauses.Selectable{
					{
						Expr: expressions.Column{Name: "column_a", From: "table_a"},
						As:   "alias_a",
					},
					{
						Expr: expressions.Column{Name: "column_b"},
					},
				},
			},
			want:    "RETURNING table_a.column_a AS alias_a, column_b",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sb, tt.args.columns); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb := tt.args.sb; sb != nil {
				if str := sb.String(); str != tt.want {
					t.Errorf("Write() got = %q, want %q", str, tt.want)
				}
			}
		})
	}

	t.Run("should return io.writer error", func(t *testing.T) {
		writer := mock.StringWriterFn(func(s string) (n int, err error) {
			return 0, errors.New("error")
		})

		if err := Write(writer, []clauses.Selectable{
			{
				Expr: expressions.Column{Name: "column_a"},
			},
		}); err == nil {
			t.Error("expected error")
		}
	})
}
