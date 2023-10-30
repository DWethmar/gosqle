package statement

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/values"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

func TestWriteInsert(t *testing.T) {
	t.Run("should write INSERT", func(t *testing.T) {
		sb := new(strings.Builder)
		if err := WriteInsert(sb, "table", []string{"column1", "column2"}); err != nil {
			t.Errorf("WriteInsert() error = %v", err)
		}

		if sb.String() != "INSERT INTO table (column1, column2)" {
			t.Errorf("WriteInsert() got = %v, want %v", sb.String(), "INSERT INTO table (column1, column2)")
		}
	})
}

func TestInsert_Write(t *testing.T) {
	type fields struct {
		ClauseWriter ClauseWriter
		table        string
		columns      []string
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
			name: "should write INSERT",
			fields: fields{
				ClauseWriter: ClauseWriter{
					clauses: map[clauses.ClauseType]clauses.Clause{
						clauses.ValuesType: values.New([]expressions.Expression{
							postgres.NewArgument("a", 1),
							postgres.NewArgument("b", 2),
						}),
					},
					order:           insertClausesOrder,
					ClauseSeparator: SpaceSeparator,
				},
				table:   "table",
				columns: []string{"column1", "column2"},
			},
			args: args{
				sw: new(strings.Builder),
			},
			want:    "INSERT INTO table (column1, column2) VALUES ($1, $2)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Insert{
				ClauseWriter: tt.fields.ClauseWriter,
				table:        tt.fields.table,
				columns:      tt.fields.columns,
			}
			if err := i.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Insert.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sw, ok := tt.args.sw.(*strings.Builder); ok {
					if sw.String() != tt.want {
						t.Errorf("Insert.Write() got = %v, want %v", sw.String(), tt.want)
					}
				} else {
					t.Errorf("expected string builder")
				}
			}
		})
	}
}

func TestNewInsert(t *testing.T) {
	type args struct {
		table   string
		columns []string
	}
	tests := []struct {
		name string
		args args
		want *Insert
	}{
		{
			name: "should create new Insert",
			args: args{
				table:   "table",
				columns: []string{"column1", "column2"},
			},
			want: &Insert{
				ClauseWriter: ClauseWriter{
					clauses:         map[clauses.ClauseType]clauses.Clause{},
					order:           insertClausesOrder,
					ClauseSeparator: SpaceSeparator,
				},
				table:   "table",
				columns: []string{"column1", "column2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInsert(tt.args.table, tt.args.columns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
