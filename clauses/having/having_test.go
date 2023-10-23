package having

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestWriteHaving(t *testing.T) {
	type args struct {
		sw    io.StringWriter
		preds []predicates.Predicate
	}
	tests := []struct {
		name        string
		args        args
		checkString bool
		want        string
		wantErr     bool
	}{
		{
			name: "should write HAVING",
			args: args{
				sw:    &strings.Builder{},
				preds: []predicates.Predicate{},
			},
			checkString: true,
			want:        "HAVING ",
			wantErr:     false,
		},
		{
			name: "should write HAVING with aggregate functions",
			args: args{
				sw: &strings.Builder{},
				preds: []predicates.Predicate{
					predicates.GT{
						Col: &expressions.Count{
							Expr: &expressions.Column{
								Name: "id",
							},
						},
						Expr: postgres.NewArgument(13, 1),
					},
					predicates.LT{
						Col: &expressions.Max{
							Expr: &expressions.Column{
								Name: "kipsate",
							},
						},
						Expr: postgres.NewArgument(9001, 2),
					},
					predicates.GTE{
						Col: &expressions.Avg{
							Expr: &expressions.Column{
								Name: "saus",
							},
						},
						Expr:  postgres.NewArgument(10000, 3),
						Logic: predicates.OR,
					},
				},
			},
			checkString: true,
			want:        "HAVING COUNT(id) > $1 AND MAX(kipsate) < $2 OR AVG(saus) >= $3",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteHaving(tt.args.sw, tt.args.preds); (err != nil) != tt.wantErr {
				t.Errorf("WriteHaving() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.checkString {
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					if got := sb.String(); got != tt.want {
						t.Errorf("WriteHaving() = %q, want %q", got, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}

func TestClause_Type(t *testing.T) {
	t.Run("should return HavingType", func(t *testing.T) {
		c := &Clause{}
		if got := c.Type(); got != clauses.HavingType {
			t.Errorf("Clause.Type() = %v, want %v", got, clauses.HavingType)
		}
	})
}

func TestClause_Write(t *testing.T) {
	type fields struct {
		Predicates []predicates.Predicate
	}
	type args struct {
		sw io.StringWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "should write HAVING",
			fields: fields{},
			args: args{
				sw: &strings.Builder{},
			},
			wantErr: false,
		},
		{
			name: "should write HAVING with aggregate functions",
			fields: fields{
				Predicates: []predicates.Predicate{
					predicates.GT{
						Col: &expressions.Count{
							Expr: &expressions.Column{
								Name: "id",
							},
						},
						Expr: postgres.NewArgument(13, 1),
					},
				},
			},
			args: args{
				sw: &strings.Builder{},
			},
			wantErr: false,
		},
		{
			name: "should return error when writing fails",
			fields: fields{
				Predicates: []predicates.Predicate{
					predicates.GT{
						Col: &expressions.Count{
							Expr: &expressions.Column{
								Name: "id",
							},
						},
						Expr: postgres.NewArgument(13, 1),
					},
				},
			},
			args: args{
				sw: mock.StringWriterFn(func(s string) (n int, err error) {
					return 0, errors.New("error")
				}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Clause{
				Predicates: tt.fields.Predicates,
			}
			if err := c.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Clause.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		predicates []predicates.Predicate
	}
	tests := []struct {
		name string
		args args
		want *Clause
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.predicates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
