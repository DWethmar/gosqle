package clauses

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
)

func TestSelectable_Write(t *testing.T) {
	type fields struct {
		Expr expressions.Expression
		As   string
	}
	type args struct {
		sw io.StringWriter
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		checkString bool
		want        string
		wantErr     bool
	}{
		{
			name: "write to string builder",
			fields: fields{
				Expr: expressions.Column{Name: "id"},
				As:   "id",
			},
			args: args{
				sw: &strings.Builder{},
			},
			checkString: true,
			want:        "id AS id",
			wantErr:     false,
		},
		{
			name: "write to string builder without as",
			fields: fields{
				Expr: expressions.Column{Name: "id"},
			},
			args: args{
				sw: &strings.Builder{},
			},
			checkString: true,
			want:        "id",
			wantErr:     false,
		},
		{
			name: "write to string builder with expression",
			fields: fields{
				Expr: expressions.Column{Name: "id"},
				As:   "id",
			},
			args: args{
				sw: &strings.Builder{},
			},
			checkString: true,
			want:        "id AS id",
			wantErr:     false,
		},
		{
			name: "write to string builder with expression without as",
			fields: fields{
				Expr: expressions.Column{Name: "id"},
			},
			args: args{
				sw: &strings.Builder{},
			},
			checkString: true,
			want:        "id",
			wantErr:     false,
		},
		{
			name: "return error when write to fails",
			fields: fields{
				Expr: expressions.Column{Name: "id"},
			},
			args: args{
				sw: mock.StringWriterFn(func(s string) (n int, err error) {
					return 0, errors.New("error")
				}),
			},
			checkString: false,
			want:        "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Selectable{
				Expr: tt.fields.Expr,
				As:   tt.fields.As,
			}
			if err := s.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Selectable.Write() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.checkString {
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					if got := sb.String(); got != tt.want {
						t.Errorf("Selectable.Write() = %q, want %q", got, tt.want)
					}
				} else {
					t.Errorf("expected string builder")
				}
			}
		})
	}
}
