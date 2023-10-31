package functions

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
)

func TestSum_Write(t *testing.T) {
	type fields struct {
		Col *expressions.Column
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
			name: "write sum",
			fields: fields{
				Col: &expressions.Column{
					Name: "test",
				},
			},
			args: args{
				sw: new(strings.Builder),
			},
			want:    "SUM(test)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sum{
				Col: tt.fields.Col,
			}

			if err := s.Write(tt.args.sw); (err != nil) != tt.wantErr {
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

func TestNewSum(t *testing.T) {
	type args struct {
		column *expressions.Column
	}
	tests := []struct {
		name string
		args args
		want *Sum
	}{
		{
			name: "new sum",
			args: args{
				column: &expressions.Column{
					Name: "test",
				},
			},
			want: &Sum{
				Col: &expressions.Column{
					Name: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSum(tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
