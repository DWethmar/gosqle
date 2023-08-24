package set

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/mock"
	"github.com/dwethmar/gosqle/postgres"
)

func TestWriteSet(t *testing.T) {
	type args struct {
		sb          *strings.Builder
		changes     []Change
		paramOffset int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "no fields supplied should return error",
			args: args{
				sb:          new(strings.Builder),
				changes:     []Change{},
				paramOffset: 0,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "set two test fields should return correct string",
			args: args{
				sb: new(strings.Builder),
				changes: []Change{
					{
						Col:  "field_a",
						Expr: postgres.NewArgument("value_a", 1),
					},
					{
						Col:  "field_b",
						Expr: postgres.NewArgument("value_a", 2),
					},
				},
			},
			want:    "SET field_a = $1, field_b = $2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sb, tt.args.changes); (err != nil) != tt.wantErr {
				t.Errorf("WriteSet() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb := tt.args.sb; sb != nil {
				if str := sb.String(); str != tt.want {
					t.Errorf("WriteSet() got = %q, want %q", str, tt.want)
				}
			}
		})
	}

	t.Run("should return io.writer error", func(t *testing.T) {
		writer := mock.StringWriterFn(func(s string) (n int, err error) {
			return 0, errors.New("error")
		})

		if err := Write(writer, []Change{
			{
				Col:  "field_a",
				Expr: postgres.NewArgument("value_a", 1),
			},
		}); err == nil {
			t.Errorf("Set() error = %v, wantErr %v", err, true)
		}
	})
}
