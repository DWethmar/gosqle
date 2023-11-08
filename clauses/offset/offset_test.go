package offset

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
		sb   *strings.Builder
		expr expressions.Expression
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should write offset mysql",
			args: args{
				sb:   &strings.Builder{},
				expr: mysql.NewArgument(1),
			},
			want:    "OFFSET ?",
			wantErr: false,
		},
		{
			name: "should write offset postgres",
			args: args{
				sb:   &strings.Builder{},
				expr: postgres.NewArgument(1, 1),
			},
			want:    "OFFSET $1",
			wantErr: false,
		},
		{
			name: "should write limit at offset 13",
			args: args{
				sb:   &strings.Builder{},
				expr: postgres.NewArgument(13, 1),
			},
			want:    "OFFSET $13",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sb, tt.args.expr); (err != nil) != tt.wantErr {
				t.Errorf("WriteOffset() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb := tt.args.sb; sb != nil {
				if str := sb.String(); str != tt.want {
					t.Errorf("WriteOffset() got = %q, want %q", str, tt.want)
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
