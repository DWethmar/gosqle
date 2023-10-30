package groupby

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
)

func TestWriteGroupByColumns(t *testing.T) {
	type args struct {
		sw      io.StringWriter
		columns []*expressions.Column
	}
	tests := []struct {
		name        string
		args        args
		checkString bool
		want        string
		wantErr     bool
	}{
		{
			name: "should write group by with one column",
			args: args{
				sw: &strings.Builder{},
				columns: []*expressions.Column{
					{Name: "id", From: "users"},
				},
			},
			checkString: true,
			want:        "GROUP BY users.id",
			wantErr:     false,
		},
		{
			name: "should write group by with multiple columns",
			args: args{
				sw: &strings.Builder{},
				columns: []*expressions.Column{
					{Name: "id", From: "users"},
					{Name: "email", From: "users"},
				},
			},
			checkString: true,
			want:        "GROUP BY users.id, users.email",
			wantErr:     false,
		},
		{
			name: "should return error when unable to write GROUP BY",
			args: args{
				sw: mock.StringWriterFn(func(s string) (n int, err error) {
					return 0, errors.New("some error")
				}),
				columns: []*expressions.Column{
					{Name: "id", From: "users"},
				},
			},
			checkString: false,
			want:        "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteGroupByColumns(tt.args.sw, tt.args.columns); (err != nil) != tt.wantErr {
				t.Errorf("WriteGroupByColumns() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.checkString {
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					if s := sb.String(); s != tt.want {
						t.Errorf("WriteGroupByColumns() = %v, want %v", s, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}

func TestClause_Type(t *testing.T) {
	t.Run("should return GroupByType", func(t *testing.T) {
		c := &Clause{}
		if got := c.Type(); got != clauses.GroupByType {
			t.Errorf("Clause.Type() = %v, want %v", got, clauses.GroupByType)
		}
	})
}

func TestClause_Write(t *testing.T) {
	type fields struct {
		Grouping Grouping
	}
	type args struct {
		sw io.StringWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should write group by with one column",
			fields: fields{
				Grouping: &ColumnGrouping{
					&expressions.Column{Name: "id", From: "users"},
					&expressions.Column{Name: "email", From: "users"},
				},
			},
			args: args{
				sw: &strings.Builder{},
			},
			wantErr: false,
		},
		{
			name: "should write group by with multiple columns",
			fields: fields{
				Grouping: &ColumnGrouping{
					&expressions.Column{Name: "id", From: "users"},
					&expressions.Column{Name: "email", From: "users"},
				},
			},
			args: args{
				sw: &strings.Builder{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Clause{
				grouping: tt.fields.Grouping,
			}
			if err := c.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Clause.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		Grouping Grouping
	}
	tests := []struct {
		name string
		args args
		want *Clause
	}{
		{
			name: "should create a new clause",
			args: args{
				Grouping: &ColumnGrouping{
					&expressions.Column{Name: "id", From: "users"},
					&expressions.Column{Name: "email", From: "users"},
				},
			},
			want: &Clause{
				grouping: &ColumnGrouping{
					&expressions.Column{Name: "id", From: "users"},
					&expressions.Column{Name: "email", From: "users"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.Grouping); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
