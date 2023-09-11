package statement

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/where"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestWriteDelete(t *testing.T) {
	t.Run("should write DELETE", func(t *testing.T) {
		sb := new(strings.Builder)
		if err := WriteDelete(sb, "table"); err != nil {
			t.Errorf("WriteDelete() error = %v", err)
		}

		if sb.String() != "DELETE FROM table" {
			t.Errorf("WriteDelete() got = %v, want %v", sb.String(), "DELETE FROM table")
		}
	})
}

func TestDelete_WriteTo(t *testing.T) {
	type fields struct {
		ClauseWriter ClauseWriter
		table        string
	}
	type args struct {
		sw io.StringWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should write DELETE",
			fields: fields{
				ClauseWriter: ClauseWriter{
					clauses: map[clauses.ClauseType]clauses.Clause{
						clauses.WhereType: where.New([]predicates.Predicate{
							predicates.EQ{
								Col:  expressions.Column{Name: "id"},
								Expr: postgres.NewArgument(1, 1),
							},
						}),
					},
					order:           deleteClausesOrder,
					ClauseSeparator: SpaceSeparator,
				},
				table: "table",
			},
			args: args{
				sw: new(strings.Builder),
			},
			want:    `DELETE FROM table WHERE id = $1`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Delete{
				ClauseWriter: tt.fields.ClauseWriter,
				table:        tt.fields.table,
			}
			if err := d.WriteTo(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Delete.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sw, ok := tt.args.sw.(*strings.Builder); ok {
					if sw.String() != tt.want {
						t.Errorf("Delete.WriteTo() got = %v, want %v", sw.String(), tt.want)
					}
				} else {
					t.Errorf("expected string builder")
				}
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
		want *Delete
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
