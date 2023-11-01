package where

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/mock"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestWhere(t *testing.T) {
	type args struct {
		sb         *strings.Builder
		conditions []logic.Logic
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "empty conditions should return error",
			args: args{
				sb:         &strings.Builder{},
				conditions: []logic.Logic{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "nil conditions should return error",
			args: args{
				sb:         &strings.Builder{},
				conditions: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "where clause with one condition",
			args: args{
				sb: &strings.Builder{},
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "id"}, postgres.NewArgument(1, 1))),
				},
			},
			want:    `WHERE id = $1`,
			wantErr: false,
		},
		{
			name: "where clause with multiple conditions",
			args: args{
				sb: &strings.Builder{},
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "id"}, postgres.NewArgument(1, 1))),
					logic.And(predicates.EQ(expressions.Column{Name: "name"}, postgres.NewArgument("piet", 2))),
				},
			},
			want:    `WHERE id = $1 AND name = $2`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteWhere(tt.args.sb, tt.args.conditions); (err != nil) != tt.wantErr {
				t.Errorf("Where() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb := tt.args.sb; sb != nil {
				if str := sb.String(); str != tt.want {
					t.Errorf("Where() got = %q, want %q", str, tt.want)
				}
			}
		})
	}

	t.Run("should return io.writer error", func(t *testing.T) {
		writer := mock.StringWriterFn(func(s string) (n int, err error) {
			return 0, errors.New("error")
		})

		if err := WriteWhere(writer, []logic.Logic{
			logic.And(predicates.EQ(expressions.Column{Name: "id"}, postgres.NewArgument(1, 1))),
			logic.And(predicates.EQ(expressions.Column{Name: "name"}, postgres.NewArgument("piet", 2))),
		}); err == nil {
			t.Error("expected error")
		}
	})
}
