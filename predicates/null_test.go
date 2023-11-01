package predicates

import (
	"reflect"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
)

func TestIsNull(t *testing.T) {
	type args struct {
		target expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should return IS NULL",
			args: args{
				target: expressions.Column{Name: "target", From: "table"},
			},
			want: &Comparison{
				Left:     expressions.Column{Name: "target", From: "table"},
				Operator: "IS",
				Right:    expressions.String("NULL"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNull(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotNull(t *testing.T) {
	type args struct {
		target expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should return IS NOT NULL",
			args: args{
				target: expressions.Column{Name: "target", From: "table"},
			},
			want: &Comparison{
				Left:     expressions.Column{Name: "target", From: "table"},
				Operator: "IS NOT",
				Right:    expressions.String("NULL"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotNull(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsNotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}
