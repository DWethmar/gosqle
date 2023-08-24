package statement

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
)

func TestSelect(t *testing.T) {
	type args struct {
		sb          *strings.Builder
		expressions []clauses.Selectable
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
				sb:          new(strings.Builder),
				expressions: []clauses.Selectable{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should render Select clause",
			args: args{
				sb: new(strings.Builder),
				expressions: []clauses.Selectable{
					{
						Expr: expressions.Column{
							From: "table_a",
							Name: "field_a",
						},
						As: "alias_a",
					},
					{
						Expr: expressions.Column{
							From: "table_b",
							Name: "field_b",
						},
						As: "alias_b",
					},
				},
			},
			want:    "SELECT table_a.field_a AS alias_a, table_b.field_b AS alias_b",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteSelect(tt.args.sb, tt.args.expressions); (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb := tt.args.sb; sb != nil {
				if str := sb.String(); str != tt.want {
					t.Errorf("Select() got = %q, want %q", str, tt.want)
				}
			}
		})
	}

	t.Run("should return io.writer error", func(t *testing.T) {
		writer := mock.StringWriterFn(func(s string) (n int, err error) {
			return 0, errors.New("error")
		})

		if err := WriteSelect(writer, []clauses.Selectable{
			{
				Expr: expressions.Column{
					From: "table_a",
					Name: "field_a",
				},
				As: "alias_a",
			},
		}); err == nil {
			t.Error("expected error")
		}
	})
}
