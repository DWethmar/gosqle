package values

import (
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

func TestWriteValues(t *testing.T) {
	type args struct {
		sb     *strings.Builder
		values []expressions.Expression
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return error when no values are given",
			args: args{
				sb:     &strings.Builder{},
				values: []expressions.Expression{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should write values",
			args: args{
				sb: &strings.Builder{},
				values: []expressions.Expression{
					postgres.NewArgument(1, "a"),
					postgres.NewArgument(2, "b"),
					postgres.NewArgument(3, "c"),
				},
			},
			want:    "VALUES ($1, $2, $3)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sb, tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("Values() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb := tt.args.sb; sb != nil {
				if str := sb.String(); str != tt.want {
					t.Errorf("Values() got = %q, want %q", str, tt.want)
				}
			}
		})
	}
}
