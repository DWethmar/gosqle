package predicates

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

func TestRange_Write(t *testing.T) {
	type fields struct {
		Target expressions.Expression
		Low    expressions.Expression
		High   expressions.Expression
	}
	type args struct {
		writer io.StringWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should write",
			fields: fields{
				Target: expressions.Column{Name: "target", From: "table"},
				Low:    postgres.NewArgument(1, 100),
				High:   postgres.NewArgument(2, 200),
			},
			args: args{
				writer: new(strings.Builder),
			},
			want: "$1 AND $2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Range{
				Low:  tt.fields.Low,
				High: tt.fields.High,
			}

			if err := r.Write(tt.args.writer); (err != nil) != tt.wantErr {
				t.Errorf("Range.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sb, ok := tt.args.writer.(*strings.Builder); ok {
					if s := sb.String(); s != tt.want {
						t.Errorf("RangeWrite() = %v, want %v", s, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		target expressions.Expression
		low    expressions.Expression
		high   expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should create range",
			args: args{
				target: expressions.Column{Name: "target", From: "table"},
				low:    postgres.NewArgument(100, 1),
				high:   postgres.NewArgument(200, 2),
			},
			want: &Comparison{
				Left:     expressions.Column{Name: "target", From: "table"},
				Operator: "BETWEEN",
				Right: &Range{
					Low:  postgres.NewArgument(100, 1),
					High: postgres.NewArgument(200, 2),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Between(tt.args.target, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Between() = %v, want %v", got, tt.want)
			}
		})
	}
}
