package gosqle

import (
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/groupby"
	"github.com/dwethmar/gosqle/clauses/join"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
	"github.com/dwethmar/gosqle/predicates"
)

func TestSelect_ToSQL(t *testing.T) {
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
				clauses.Selectable{Expr: expressions.NewColumn("username").SetFrom("u")},
				clauses.Selectable{Expr: expressions.NewColumn("country").SetFrom("u"), As: "c"},
			).From(expressions.Table{
				Name:  "users",
				Alias: "u",
			}),
			want:    "SELECT u.username, u.country AS c FROM users u;",
			wantErr: false,
		},
		{
			name: "select all",
			sel: NewSelect(
				clauses.Selectable{Expr: expressions.Column{Name: "*"}},
			).From(expressions.Table{
				Name: "users",
			}),
			want:    "SELECT * FROM users;",
			wantErr: false,
		},
		{
			name: "select with joins",
			sel: NewSelect(
				clauses.Selectable{Expr: expressions.Column{Name: "*", From: "users"}},
				clauses.Selectable{Expr: expressions.Column{Name: "*", From: "companies"}},
			).From(expressions.Table{
				Name: "users",
			}).Join(
				join.Options{
					Type: join.LeftJoin,
					From: "companies",
					Match: &join.JoinOn{
						Predicates: []predicates.Predicate{
							predicates.EQ{
								Col:  expressions.NewColumn("id").SetFrom("companies"),
								Expr: expressions.NewColumn("company_id").SetFrom("users"),
							},
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
				clauses.Selectable{Expr: expressions.Column{Name: "*"}},
			).From(expressions.Table{
				Name: "users",
			}).Where(
				predicates.EQ{
					Col:  expressions.NewColumn("username"),
					Expr: postgres.NewArgument("david", 1),
				},
				predicates.EQ{
					Col:   expressions.NewColumn("email"),
					Expr:  postgres.NewArgument("test@test.com", 2),
					Logic: predicates.OR,
				},
			),
			want:    "SELECT * FROM users WHERE username = $1 OR email = $2;",
			wantErr: false,
		},
		{
			name: "select order by",
			sel: NewSelect(
				clauses.Selectable{Expr: expressions.Column{Name: "*"}},
			).From(expressions.Table{
				Name: "users",
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
				clauses.Selectable{Expr: expressions.Column{Name: "*"}},
			).From(expressions.Table{
				Name: "users",
			}).Limit(postgres.NewArgument(12, 1)),
			want:    "SELECT * FROM users LIMIT $1;",
			wantErr: false,
		},
		{
			name: "select offset",
			sel: NewSelect(
				clauses.Selectable{Expr: expressions.Column{Name: "*"}},
			).From(expressions.Table{
				Name: "users",
			}).Offset(postgres.NewArgument(100, 1)),
			want:    "SELECT * FROM users OFFSET $1;",
			wantErr: false,
		},
		{
			name: "select group by columns",
			sel: NewSelect(
				clauses.Selectable{Expr: expressions.Column{Name: "username"}},
			).From(expressions.Table{
				Name: "users",
			}).GroupBy(
				&groupby.ColumnGrouping{
					expressions.NewColumn("username"),
					expressions.NewColumn("email"),
				},
			),
			want:    "SELECT username FROM users GROUP BY username, email;",
			wantErr: false,
		},
		{
			name: "select group by columns with having",
			sel: NewSelect(
				clauses.Selectable{
					Expr: expressions.Column{Name: "username"},
				},
			).From(expressions.Table{
				Name: "users",
			}).GroupBy(
				&groupby.ColumnGrouping{
					expressions.NewColumn("username"),
					expressions.NewColumn("email"),
				},
			).Having(
				predicates.GT{
					Col:  expressions.NewMax(expressions.NewColumn("id")),
					Expr: postgres.NewArgument(12, 1),
				},
				predicates.LT{
					Col:   expressions.NewMax(expressions.NewColumn("id")),
					Expr:  postgres.NewArgument(12, 2),
					Logic: predicates.OR,
				},
			),
			want:    "SELECT username FROM users GROUP BY username, email HAVING MAX(id) > $1 OR MAX(id) < $2;",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := new(strings.Builder)
			err := tt.sel.WriteTo(sb)

			if (err != nil) != tt.wantErr {
				t.Errorf("Select.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if query := sb.String(); query != tt.want {
				t.Errorf("Select.WriteTo() query = %q, wantQuery %q", query, tt.want)
			}
		})
	}
}
