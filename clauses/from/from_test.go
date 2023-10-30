package from

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
)

func TestWrite(t *testing.T) {
	type args struct {
		sw   io.StringWriter
		from From
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should write From",
			args: args{
				sw: &strings.Builder{},
				from: From{
					Expr: expressions.String("table"),
				},
			},
			want:    "FROM table",
			wantErr: false,
		},
		{
			name: "should write From with alias",
			args: args{
				sw: &strings.Builder{},
				from: From{
					Expr: expressions.String("table"),
					As:   "alias",
				},
			},
			want:    "FROM table AS alias",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sw, tt.args.from); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.want != "" {
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					if got := sb.String(); got != tt.want {
						t.Errorf("Write() = %q, want %q", got, tt.want)
					}
				} else {
					t.Errorf("expected string builder")
				}
			}
		})
	}
}

func TestFrom_Type(t *testing.T) {
	t.Run("should return FromClauseType", func(t *testing.T) {
		f := &Clause{}
		if got := f.Type(); got != clauses.FromType {
			t.Errorf("From.Type() = %v, want %v", got, clauses.FromType)
		}
	})
}

func TestFrom_Write(t *testing.T) {
	type fields struct {
		Expression expressions.Expression
	}
	type args struct {
		sb *strings.Builder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "should write From",
			fields:  fields{Expression: expressions.String("table")},
			args:    args{sb: new(strings.Builder)},
			want:    "FROM table",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Clause{
				from: From{
					Expr: tt.fields.Expression,
				},
			}
			if err := f.Write(tt.args.sb); (err != nil) != tt.wantErr {
				t.Errorf("From.Write() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				got := tt.args.sb.String()
				if got != tt.want {
					t.Errorf("From.Write() = %q, want %q", got, tt.want)
				}
			}
		})
	}
}

func TestNewFrom(t *testing.T) {
	type args struct {
		from From
	}
	tests := []struct {
		name string
		args args
		want *Clause
	}{
		{
			name: "should create new From",
			args: args{
				from: From{
					Expr: expressions.String("table"),
					As:   "",
				},
			},
			want: &Clause{
				from: From{
					Expr: expressions.String("table"),
					As:   "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
