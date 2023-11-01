package alias

import (
	"io"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
)

func TestAlias_Write(t *testing.T) {
	type fields struct {
		Expr expressions.Expression
		As   string
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
			name: "should write alias",
			fields: fields{
				Expr: expressions.String("table"),
				As:   "alias",
			},
			args: args{
				sw: &strings.Builder{},
			},
			want:    "table AS alias",
			wantErr: false,
		},
		{
			name: "should write alias without AS",
			fields: fields{
				Expr: expressions.String("table"),
				As:   "",
			},
			args: args{
				sw: &strings.Builder{},
			},
			want:    "table",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Alias{
				Expr: tt.fields.Expr,
				As:   tt.fields.As,
			}

			if err := a.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Alias.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					if got := sb.String(); got != tt.want {
						t.Errorf("Alias.Write() = %q, want %q", got, tt.want)
					}
				} else {
					t.Errorf("expected string builder")
				}
			}
		})
	}
}
