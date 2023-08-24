package limit

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
	"github.com/dwethmar/gosqle/mysql"
	"github.com/dwethmar/gosqle/postgres"
)

func TestWrite(t *testing.T) {
	type args struct {
		sb  *strings.Builder
		arg expressions.Expression
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should write limit mysql",
			args: args{
				sb:  &strings.Builder{},
				arg: mysql.NewArgument(100),
			},
			want:    "LIMIT ?",
			wantErr: false,
		},
		{
			name: "should write limit postgres",
			args: args{
				sb:  &strings.Builder{},
				arg: postgres.NewArgument(100, 1),
			},
			want:    "LIMIT $1",
			wantErr: false,
		},
		{
			name: "should write limit at index 13",
			args: args{
				sb:  &strings.Builder{},
				arg: postgres.NewArgument(100, 13),
			},
			want:    "LIMIT $13",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sb, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb := tt.args.sb; sb != nil {
				if str := sb.String(); str != tt.want {
					t.Errorf("Write() got = %q, want %q", str, tt.want)
				}
			}
		})
	}

	t.Run("should return io.writer error", func(t *testing.T) {
		writer := mock.StringWriterFn(func(s string) (n int, err error) {
			return 0, errors.New("error")
		})

		if err := Write(writer, mysql.NewArgument(1)); err == nil {
			t.Error("expected error")
		}
	})
}
