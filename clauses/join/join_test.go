package join

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/mock"
	"github.com/dwethmar/gosqle/predicates"
)

func TestWriteJoin(t *testing.T) {
	type args struct {
		sb       *strings.Builder
		joinType Type
		from     string
		Match    Matcher
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "inner join",
			args: args{
				sb:       new(strings.Builder),
				joinType: InnerJoin,
				from:     "table",
				Match: &On{
					Conditions: []logic.Logic{
						logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
						logic.And(predicates.LT(expressions.Column{Name: "field_c", From: "table"}, expressions.Column{Name: "field_d", From: "other_table"})),
					},
				},
			},
			want:    "INNER JOIN table ON table.field_a = other_table.field_b AND table.field_c < other_table.field_d",
			wantErr: false,
		},
		{
			name: "left join",
			args: args{
				sb:       new(strings.Builder),
				joinType: LeftJoin,
				from:     "table",
				Match: &On{
					Conditions: []logic.Logic{
						logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
						logic.And(predicates.LT(expressions.Column{Name: "field_c", From: "table"}, expressions.Column{Name: "field_d", From: "other_table"})),
					},
				},
			},
			want:    "LEFT JOIN table ON table.field_a = other_table.field_b AND table.field_c < other_table.field_d",
			wantErr: false,
		},
		{
			name: "right join",
			args: args{
				sb:       new(strings.Builder),
				joinType: RightJoin,
				from:     "table",
				Match: &On{
					Conditions: []logic.Logic{
						logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
						logic.And(predicates.LT(expressions.Column{Name: "field_c", From: "table"}, expressions.Column{Name: "field_d", From: "other_table"})),
					},
				},
			},
			want:    "RIGHT JOIN table ON table.field_a = other_table.field_b AND table.field_c < other_table.field_d",
			wantErr: false,
		},
		{
			name: "full join",
			args: args{
				sb:       new(strings.Builder),
				joinType: FullJoin,
				from:     "table",
				Match: &On{
					Conditions: []logic.Logic{
						logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
						logic.And(predicates.LT(expressions.Column{Name: "field_c", From: "table"}, expressions.Column{Name: "field_d", From: "other_table"})),
					},
				},
			},
			want:    "FULL JOIN table ON table.field_a = other_table.field_b AND table.field_c < other_table.field_d",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.sb, tt.args.joinType, tt.args.from, tt.args.Match); (err != nil) != tt.wantErr {
				t.Errorf("WriteJoin() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				got := tt.args.sb.String()
				if got != tt.want {
					t.Errorf("WriteJoin() = %q, want %q", got, tt.want)
				}
			}
		})
	}
}

func TestJoin_Type(t *testing.T) {
	t.Run("should return JoinClauseType", func(t *testing.T) {
		j := &Clause{}
		if got := j.Type(); got != clauses.JoinType {
			t.Errorf("Join.Type() = %v, want %v", got, clauses.JoinType)
		}
	})
}

func TestJoin_Write(t *testing.T) {
	type fields struct {
		joins []Options
	}
	type args struct {
		sb *strings.Builder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should write join with ON match",
			fields: fields{
				joins: []Options{
					{
						Type: InnerJoin,
						From: "table",
						Match: &On{
							Conditions: []logic.Logic{
								logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
							},
						},
					},
				},
			},
			args:    args{sb: new(strings.Builder)},
			want:    "INNER JOIN table ON table.field_a = other_table.field_b",
			wantErr: false,
		},
		{
			name: "should write join with USING match",
			fields: fields{
				joins: []Options{
					{
						Type: InnerJoin,
						From: "table",
						Match: &Using{
							Uses: []string{"field_a", "field_b"},
						},
					},
				},
			},
			args:    args{sb: new(strings.Builder)},
			want:    "INNER JOIN table USING (field_a, field_b)",
			wantErr: false,
		},
		{
			name: "should write multiple joins",
			fields: fields{
				joins: []Options{
					{
						Type: InnerJoin,
						From: "table_b",
						Match: &On{
							Conditions: []logic.Logic{
								logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table_a"}, expressions.Column{Name: "field_b", From: "table_b"})),
							},
						},
					},
					{
						Type: LeftJoin,
						From: "table_c",
						Match: &On{
							Conditions: []logic.Logic{
								logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table_a"}, expressions.Column{Name: "field_b", From: "table_c"})),
							},
						},
					},
				},
			},
			args:    args{sb: new(strings.Builder)},
			want:    "INNER JOIN table_b ON table_a.field_a = table_b.field_b LEFT JOIN table_c ON table_a.field_a = table_c.field_b",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Clause{
				joins: tt.fields.joins,
			}
			if err := j.Write(tt.args.sb); (err != nil) != tt.wantErr {
				t.Errorf("Join.Write() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				got := tt.args.sb.String()
				if got != tt.want {
					t.Errorf("Join.Write() = %q, want %q", got, tt.want)
				}
			}
		})
	}
}

func TestNewJoin(t *testing.T) {
	type args struct {
		j []Options
	}
	tests := []struct {
		name string
		args args
		want *Clause
	}{
		{
			name: "should create new Join",
			args: args{
				j: []Options{
					{
						Type: InnerJoin,
						From: "table",
						Match: &On{
							Conditions: []logic.Logic{
								logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
							},
						},
					},
				},
			},
			want: &Clause{
				joins: []Options{
					{
						Type: InnerJoin,
						From: "table",
						Match: &On{
							Conditions: []logic.Logic{
								logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinOn_Write(t *testing.T) {
	type fields struct {
		conditions []logic.Logic
	}
	type args struct {
		sw io.StringWriter
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		checkString bool
		want        string
		wantErr     bool
	}{
		{
			name: "should return error when writer is nil",
			fields: fields{
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
				},
			},
			args:    args{sw: nil},
			want:    "",
			wantErr: true,
		},
		{
			name: "should return error when no conditions",
			fields: fields{
				conditions: []logic.Logic{},
			},
			args:        args{sw: new(strings.Builder)},
			checkString: false,
			want:        "",
			wantErr:     true,
		},
		{
			name: "should write join with ON match",
			fields: fields{
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
				},
			},
			args:        args{sw: new(strings.Builder)},
			checkString: true,
			want:        "ON table.field_a = other_table.field_b",
			wantErr:     false,
		},
		{
			name: "should return error when string writer returns error",
			fields: fields{
				conditions: []logic.Logic{
					logic.And(predicates.EQ(expressions.Column{Name: "field_a", From: "table"}, expressions.Column{Name: "field_b", From: "other_table"})),
				},
			},
			args: args{
				sw: mock.StringWriterFn(func(s string) (int, error) {
					return 0, errors.New("error")
				}),
			},
			checkString: false,
			want:        "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &On{
				Conditions: tt.fields.conditions,
			}
			if err := j.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("JoinOn.Write() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.checkString {
				// Check if we can get the string from the string builder.
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					got := sb.String()
					if got != tt.want {
						t.Errorf("JoinOn.Write() = %q, want %q", got, tt.want)
					}
				} else {
					t.Errorf("could not get string from string builder")
				}
			}
		})
	}
}

func TestJoinUsing_Write(t *testing.T) {
	type fields struct {
		Uses []string
	}
	type args struct {
		sw io.StringWriter
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		checkString bool
		want        string
		wantErr     bool
	}{
		{
			name: "should write join with USING match",
			fields: fields{
				Uses: []string{"field_a", "field_b"},
			},
			args: args{
				sw: new(strings.Builder),
			},
			checkString: true,
			want:        "USING (field_a, field_b)",
			wantErr:     false,
		},
		{
			name: "should return error when string writer returns error",
			fields: fields{
				Uses: []string{"field_a", "field_b"},
			},
			args: args{
				sw: mock.StringWriterFn(func(s string) (int, error) {
					return 0, errors.New("error")
				}),
			},
			checkString: false,
			want:        "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Using{
				Uses: tt.fields.Uses,
			}
			if err := j.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("JoinUsing.Write() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.checkString {
				// Check if we can get the string from the string builder.
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					got := sb.String()
					if got != tt.want {
						t.Errorf("JoinUsing.Write() = %q, want %q", got, tt.want)
					}
				} else {
					t.Errorf("could not get string from string builder")
				}
			}
		})
	}
}
