package having

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestWriteHaving(t *testing.T) {
	type args struct {
		sw         io.StringWriter
		conditions []logic.Logic
	}
	tests := []struct {
		name        string
		args        args
		checkString bool
		want        string
		wantErr     bool
	}{
		{
			name: "should return error when writer is nil",
			args: args{
				sw: nil,
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "id"}, postgres.NewArgument(1, 1))),
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should return error when conditions is nil",
			args: args{
				sw:         new(strings.Builder),
				conditions: nil,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteHaving(tt.args.sw, tt.args.conditions); (err != nil) != tt.wantErr {
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
		conditions []logic.Logic
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
			name: "should return error when writer is nil",
			fields: fields{
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "id"}, postgres.NewArgument(1, 1))),
				},
			},
			args: args{
				sw: nil,
			},
			wantErr: true,
		},
		{
			name: "should return error when conditions is nil",
			fields: fields{
				conditions: nil,
			},
			args: args{
				sw: new(strings.Builder),
			},
			wantErr: true,
		},
		{
			name: "should write",
			fields: fields{
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "id"}, postgres.NewArgument(1, 1))),
				},
			},
			args: args{
				sw: new(strings.Builder),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Clause{
				conditions: tt.fields.conditions,
			}
			if err := c.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Clause.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		conditions []logic.Logic
	}
	tests := []struct {
		name string
		args args
		want *Clause
	}{
		{
			name: "should create new clause",
			args: args{
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "id"}, postgres.NewArgument(1, 1))),
				},
			},
			want: &Clause{
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "id"}, postgres.NewArgument(1, 1))),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.conditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
