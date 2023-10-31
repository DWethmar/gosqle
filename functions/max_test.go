package functions

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
)

func TestMax_Write(t *testing.T) {
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
			name: "write max",
			fields: fields{
				Col: &expressions.Column{
					Name: "test",
				},
			},
			args: args{
				sw: new(strings.Builder),
			},
			want:    "MAX(test)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Max{
				Col: tt.fields.Col,
			}

			if err := m.Write(tt.args.sw); (err != nil) != tt.wantErr {
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

func TestNewMax(t *testing.T) {
	type args struct {
		column *expressions.Column
	}
	tests := []struct {
		name string
		args args
		want *Max
	}{
		{
			name: "new max",
			args: args{
				column: &expressions.Column{
					Name: "test",
				},
			},
			want: &Max{
				Col: &expressions.Column{
					Name: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMax(tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
