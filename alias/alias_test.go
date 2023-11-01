package alias

import (
	"io"
	"reflect"
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

func TestAlias_SetAs(t *testing.T) {
	type fields struct {
		Expr expressions.Expression
		As   string
	}
	type args struct {
		as string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Alias
	}{
		{
			name: "should set alias",
			fields: fields{
				Expr: expressions.String("table"),
				As:   "",
			},
			args: args{
				as: "alias",
			},
			want: &Alias{
				Expr: expressions.String("table"),
				As:   "alias",
			},
		},
		{
			name: "should overwrite alias",
			fields: fields{
				Expr: expressions.String("table"),
				As:   "alias",
			},
			args: args{
				as: "alias2",
			},
			want: &Alias{
				Expr: expressions.String("table"),
				As:   "alias2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Alias{
				Expr: tt.fields.Expr,
				As:   tt.fields.As,
			}
			if got := a.SetAs(tt.args.as); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Alias.SetAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		expr expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Alias
	}{
		{
			name: "should create alias",
			args: args{
				expr: expressions.String("table"),
			},
			want: &Alias{
				Expr: expressions.String("table"),
				As:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.expr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
