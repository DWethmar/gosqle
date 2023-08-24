package orderby

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
)

func TestWrite(t *testing.T) {
	type args struct {
		sb      *strings.Builder
		sorting []Sort
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return error when no fields are supplied",
			args: args{
				sb:      new(strings.Builder),
				sorting: []Sort{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should render order by clause",
			args: args{
				sb: new(strings.Builder),
				sorting: []Sort{
					{
						Column: &expressions.Column{
							Name: "field_a",
						},
						Direction: ASC,
					},
					{
						Column: &expressions.Column{
							Name: "field_b",
						},
						Direction: DESC,
					},
				},
			},
			want:    "ORDER BY field_a ASC, field_b DESC",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sb, tt.args.sorting); (err != nil) != tt.wantErr {
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

		if err := Write(writer, []Sort{
			{
				Column: &expressions.Column{
					Name: "field_a",
				},
				Direction: ASC,
			},
			{
				Column: &expressions.Column{
					Name: "field_b",
				},
				Direction: DESC,
			},
		}); err == nil {
			t.Error("expected error")
		}
	})
}
