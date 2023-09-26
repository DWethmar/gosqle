package statement

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/from"
	"github.com/dwethmar/gosqle/expressions"
)

func TestWriteSelect(t *testing.T) {
	t.Run("should write SELECT", func(t *testing.T) {
		sb := new(strings.Builder)
		if err := WriteSelect(sb, []clauses.Selectable{
			{
				Expr: expressions.Column{Name: "column1"},
			},
		}); err != nil {
			t.Errorf("WriteSelect() error = %v", err)
		}

		if sb.String() != "SELECT column1" {
			t.Errorf("WriteSelect() got = %v, want %v", sb.String(), "SELECT column1")
		}
	})
}

func TestSelect_WriteTo(t *testing.T) {
	type fields struct {
		ClauseWriter  ClauseWriter
		selectColumns []clauses.Selectable
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
			name: "should write SELECT",
			fields: fields{
				ClauseWriter: ClauseWriter{
					clauses: map[clauses.ClauseType]clauses.Clause{
						clauses.FromType: from.New(from.From{
							Expr: from.Table("table"),
						}),
					},
					order:           selectClausesOrder,
					ClauseSeparator: SpaceSeparator,
				},
				selectColumns: []clauses.Selectable{{
					Expr: expressions.Column{Name: "column1"},
				}},
			},
			args: args{
				sw: new(strings.Builder),
			},
			want:    "SELECT column1 FROM table",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Select{
				ClauseWriter:  tt.fields.ClauseWriter,
				selectColumns: tt.fields.selectColumns,
			}
			if err := s.WriteTo(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Select.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sw, ok := tt.args.sw.(*strings.Builder); ok {
					if sw.String() != tt.want {
						t.Errorf("Select.WriteTo() got = %v, want %v", sw.String(), tt.want)
					}
				} else {
					t.Errorf("expected string builder")
				}
			}
		})
	}
}

func TestNewSelect(t *testing.T) {
	type args struct {
		selectColumns []clauses.Selectable
	}
	tests := []struct {
		name string
		args args
		want *Select
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSelect(tt.args.selectColumns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
