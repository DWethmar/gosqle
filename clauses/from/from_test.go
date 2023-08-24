package from

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
)

func TestWrite(t *testing.T) {
	type args struct {
		sb    *strings.Builder
		table expressions.Table
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
				sb:    &strings.Builder{},
				table: expressions.Table{Name: "table"},
			},
			want:    "FROM table",
			wantErr: false,
		},
		{
			name: "should write From with alias",
			args: args{
				sb:    &strings.Builder{},
				table: expressions.Table{Name: "table", Alias: "alias"},
			},
			want:    "FROM table alias",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sb, tt.args.table); (err != nil) != tt.wantErr {
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

		err := Write(writer, expressions.Table{
			Name: "table", Alias: "alias"},
		)

		if err == nil {
			t.Error("expected error")
		}
	})
}

func TestFrom_Type(t *testing.T) {
	t.Run("should return FromClauseType", func(t *testing.T) {
		f := &Clause{}
		if got := f.Type(); got != clauses.FromType {
			t.Errorf("From.Type() = %v, want %v", got, clauses.FromType)
		}
	})
}

func TestFrom_WriteTo(t *testing.T) {
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
			fields:  fields{Expression: expressions.Table{Name: "table"}},
			args:    args{sb: new(strings.Builder)},
			want:    "FROM table",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Clause{
				Expression: tt.fields.Expression,
			}
			if err := f.WriteTo(tt.args.sb); (err != nil) != tt.wantErr {
				t.Errorf("From.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				got := tt.args.sb.String()
				if got != tt.want {
					t.Errorf("From.WriteTo() = %q, want %q", got, tt.want)
				}
			}
		})
	}
}

func TestNewFrom(t *testing.T) {
	type args struct {
		expr expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Clause
	}{
		{
			name: "should create new From",
			args: args{expr: expressions.Table{Name: "table"}},
			want: &Clause{Expression: expressions.Table{Name: "table"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.expr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
