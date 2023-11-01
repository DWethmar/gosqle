package logic

import (
	"io"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestWhere(t *testing.T) {
	type args struct {
		sw    io.StringWriter
		logic []Logic
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should write nothing when logic is empty",
			args: args{
				sw:    new(strings.Builder),
				logic: []Logic{},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "should write nothing when logic is nil",
			args: args{
				sw:    new(strings.Builder),
				logic: nil,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "should return error when writer is nil",
			args: args{
				sw:    nil,
				logic: []Logic{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should write one logic",
			args: args{
				sw: new(strings.Builder),
				logic: []Logic{
					{
						Operator: AndOperator,
						Condition: &predicates.Comparison{
							Left:     expressions.Column{Name: "id"},
							Operator: "=",
							Right:    postgres.NewArgument(1, 1),
						},
					},
				},
			},
			want:    "id = $1",
			wantErr: false,
		},
		{
			name: "should write multiple logic with AND",
			args: args{
				sw: new(strings.Builder),
				logic: []Logic{
					{
						Operator: AndOperator,
						Condition: &predicates.Comparison{
							Left:     expressions.Column{Name: "id"},
							Operator: "=",
							Right:    postgres.NewArgument(1, 1),
						},
					},
					{
						Operator: AndOperator,
						Condition: &predicates.Comparison{
							Left:     expressions.Column{Name: "name"},
							Operator: "=",
							Right:    postgres.NewArgument("test", 2),
						},
					},
				},
			},
			want:    "id = $1 AND name = $2",
			wantErr: false,
		},
		{
			name: "should write multiple logic with OR",
			args: args{
				sw: new(strings.Builder),
				logic: []Logic{
					{
						Operator: OrOperator,
						Condition: &predicates.Comparison{
							Left:     expressions.Column{Name: "id"},
							Operator: "=",
							Right:    postgres.NewArgument(1, 1),
						},
					},
					{
						Operator: OrOperator,
						Condition: &predicates.Comparison{
							Left:     expressions.Column{Name: "name"},
							Operator: "=",
							Right:    postgres.NewArgument("test", 2),
						},
					},
				},
			},
			want:    "id = $1 OR name = $2",
			wantErr: false,
		},
		{
			name: "should write multiple logic with AND and OR",
			args: args{
				sw: new(strings.Builder),
				logic: []Logic{
					{
						Operator: AndOperator,
						Condition: &predicates.Comparison{
							Left:     expressions.Column{Name: "id"},
							Operator: "=",
							Right:    postgres.NewArgument(1, 1),
						},
					},
					{
						Operator: OrOperator,
						Condition: &predicates.Comparison{
							Left:     expressions.Column{Name: "name"},
							Operator: "=",
							Right:    postgres.NewArgument("test", 2),
						},
					},
					{
						Operator: AndOperator,
						Condition: &predicates.Comparison{
							Left:     expressions.Column{Name: "age"},
							Operator: "=",
							Right:    postgres.NewArgument(85, 2),
						},
					},
				},
			},
			want:    "id = $1 OR name = $2 AND age = $2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Where(tt.args.sw, tt.args.logic); (err != nil) != tt.wantErr {
				t.Errorf("Where() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					if s := sb.String(); s != tt.want {
						t.Errorf("Where() = %v, want %v", s, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}
