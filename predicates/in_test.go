package predicates

import (
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
)

func TestIN(t *testing.T) {
	type args struct {
		target expressions.Expression
		expr   []expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should create IN predicate",
			args: args{
				target: expressions.String("column"),
				expr: []expressions.Expression{
					expressions.String("1"),
					expressions.String("2"),
				},
			},
			want: "column IN (1, 2)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := new(strings.Builder)

			if err := IN(tt.args.target, tt.args.expr...).Write(sw); err != nil {
				t.Errorf("IN() error = %v", err)
			}

			if sw.String() != tt.want {
				t.Errorf("IN() got = %v, want %v", sw.String(), tt.want)
			}
		})
	}
}
