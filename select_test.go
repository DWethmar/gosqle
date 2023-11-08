package gosqle

import (
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/clauses/groupby"
	"github.com/dwethmar/gosqle/clauses/join"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/functions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestSelect_Write(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		sel     *Select
		want    string
		wantErr bool
	}{
		{
			name: "select columns",
			sel: NewSelect(
				&alias.Alias{Expr: expressions.Column{Name: "username", From: "u"}},
				&alias.Alias{Expr: expressions.Column{Name: "country", From: "u"}, As: "c"},
			).From(&alias.Alias{
				Expr: expressions.String("users"),
				As:   "u",
			}),
			want:    "SELECT u.username, u.country AS c FROM users AS u;",
			wantErr: false,
		},
		{
			name: "select all",
			sel: NewSelect(
				&alias.Alias{Expr: expressions.Column{Name: "*"}},
			).From(&alias.Alias{
				Expr: expressions.String("users"),
			}),
			want:    "SELECT * FROM users;",
			wantErr: false,
		},
		{
			name: "select with joins",
			sel: NewSelect(
				&alias.Alias{Expr: expressions.Column{Name: "*", From: "users"}},
				&alias.Alias{Expr: expressions.Column{Name: "*", From: "companies"}},
			).From(&alias.Alias{
				Expr: expressions.String("users"),
			}).Join(
				join.Options{
					Type: join.LeftJoin,
					From: "companies",
					Match: &join.On{
						Conditions: []logic.Logic{
							logic.And(
								predicates.EQ(
									expressions.Column{Name: "id", From: "companies"},
									expressions.Column{Name: "company_id", From: "users"},
								),
							),
						},
					},
				},
			),
			want:    "SELECT users.*, companies.* FROM users LEFT JOIN companies ON companies.id = users.company_id;",
			wantErr: false,
		},
		{
			name: "select where",
			sel: NewSelect(
				&alias.Alias{Expr: expressions.Column{Name: "*"}},
			).From(&alias.Alias{
				Expr: expressions.String("users"),
			}).Where(
				logic.And(
					predicates.EQ(expressions.Column{Name: "username"}, postgres.NewArgument(1, 1)),
				),
				logic.Or(
					predicates.EQ(expressions.Column{Name: "email"}, postgres.NewArgument(2, 2)),
				),
			),
			want:    "SELECT * FROM users WHERE username = $1 OR email = $2;",
			wantErr: false,
		},
		{
			name: "select order by",
			sel: NewSelect(
				&alias.Alias{Expr: expressions.Column{Name: "*"}},
			).From(&alias.Alias{
				Expr: expressions.String("users"),
			}).OrderBy(
				orderby.Sort{Column: &expressions.Column{Name: "username"}, Direction: orderby.ASC},
				orderby.Sort{Column: &expressions.Column{Name: "age"}, Direction: orderby.DESC},
			),
			want:    "SELECT * FROM users ORDER BY username ASC, age DESC;",
			wantErr: false,
		},
		{
			name: "select limit",
			sel: NewSelect(
				&alias.Alias{Expr: expressions.Column{Name: "*"}},
			).From(&alias.Alias{
				Expr: expressions.String("users"),
			}).Limit(postgres.NewArgument(1, 12)),
			want:    "SELECT * FROM users LIMIT $1;",
			wantErr: false,
		},
		{
			name: "select offset",
			sel: NewSelect(
				&alias.Alias{Expr: expressions.Column{Name: "*"}},
			).From(&alias.Alias{
				Expr: expressions.String("users"),
			}).Offset(postgres.NewArgument(1, 100)),
			want:    "SELECT * FROM users OFFSET $1;",
			wantErr: false,
		},
		{
			name: "select group by columns",
			sel: NewSelect(
				&alias.Alias{Expr: expressions.Column{Name: "username"}},
			).From(&alias.Alias{
				Expr: expressions.String("users"),
			}).GroupBy(
				&groupby.ColumnGrouping{
					&expressions.Column{Name: "username"},
					&expressions.Column{Name: "email"},
				},
			),
			want:    "SELECT username FROM users GROUP BY username, email;",
			wantErr: false,
		},
		{
			name: "select group by columns with having",
			sel: NewSelect(
				alias.New(expressions.Column{Name: "username"}),
			).From(&alias.Alias{
				Expr: expressions.String("users"),
			}).GroupBy(
				&groupby.ColumnGrouping{
					&expressions.Column{Name: "username"},
					&expressions.Column{Name: "email"},
				},
			).Having(
				logic.And(
					predicates.GT(functions.NewMax(&expressions.Column{Name: "id"}), postgres.NewArgument(1, 1)),
				),
				logic.Or(
					predicates.LT(functions.NewMax(&expressions.Column{Name: "id"}), postgres.NewArgument(2, 2)),
				),
			),
			want:    "SELECT username FROM users GROUP BY username, email HAVING MAX(id) > $1 OR MAX(id) < $2;",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := new(strings.Builder)
			err := tt.sel.Write(sb)

			if (err != nil) != tt.wantErr {
				t.Errorf("Select.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if query := sb.String(); query != tt.want {
				t.Errorf("Select.Write() query = %q, wantQuery %q", query, tt.want)
			}
		})
	}
}
