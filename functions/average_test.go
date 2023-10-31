package functions

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
)

func TestAvg_Write(t *testing.T) {
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
			name: "write avg",
			fields: fields{
				Col: &expressions.Column{
					Name: "test",
				},
			},
			args: args{
				sw: new(strings.Builder),
			},
			want:    "AVG(test)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Avg{
				Col: tt.fields.Col,
			}

			if err := a.Write(tt.args.sw); (err != nil) != tt.wantErr {
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

func TestNewAvg(t *testing.T) {
	type args struct {
		column *expressions.Column
	}
	tests := []struct {
		name string
		args args
		want *Avg
	}{
		{
			name: "new avg",
			args: args{
				column: &expressions.Column{
					Name: "test",
				},
			},
			want: &Avg{
				Col: &expressions.Column{
					Name: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAvg(tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}
