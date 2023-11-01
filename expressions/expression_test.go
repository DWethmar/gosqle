package expressions

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

type testExpression struct {
	V string
}

func (e testExpression) Write(writer io.StringWriter) error {
	if _, err := writer.WriteString(e.V); err != nil {
		return fmt.Errorf("error writing test expression: %v", err)
	}

	return nil
}

func TestList_Write(t *testing.T) {
	type args struct {
		writer io.StringWriter
	}
	tests := []struct {
		name    string
		e       List
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "empty list",
			e:    List{},
			args: args{
				writer: new(strings.Builder),
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "single expression",
			e: List{
				&testExpression{V: "1"},
			},
			args: args{
				writer: new(strings.Builder),
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "multiple expressions",
			e: List{
				&testExpression{V: "1"},
				&testExpression{V: "2"},
			},
			args: args{
				writer: new(strings.Builder),
			},
			want: "1, 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Write(tt.args.writer); (err != nil) != tt.wantErr {
				t.Errorf("List.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sb, ok := tt.args.writer.(*strings.Builder); ok {
					got := sb.String()
					if got != tt.want {
						t.Errorf("List.Write() got = %v, want %v", got, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}

func TestExpressionFunc_Write(t *testing.T) {
	type args struct {
		writer io.StringWriter
	}
	tests := []struct {
		name    string
		e       ExpressionFunc
		args    args
		wantErr bool
	}{
		{
			name: "should write",
			e: ExpressionFunc(func(writer io.StringWriter) error {
				if _, err := writer.WriteString("test"); err != nil {
					return fmt.Errorf("error writing test expression: %v", err)
				}

				return nil
			}),
			args: args{
				writer: new(strings.Builder),
			},
			wantErr: false,
		},
		{
			name: "should error",
			e: ExpressionFunc(func(writer io.StringWriter) error {
				return fmt.Errorf("error")
			}),
			args: args{
				writer: new(strings.Builder),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.Write(tt.args.writer); (err != nil) != tt.wantErr {
				t.Errorf("ExpressionFunc.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWrapInParenthesis(t *testing.T) {
	type args struct {
		sw io.StringWriter
		e  Expression
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should wrap",
			args: args{
				sw: &strings.Builder{},
				e:  testExpression{V: "1"},
			},
			want:    "(1)",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WrapInParenthesis(tt.args.sw, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("WrapInParenthesis() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					got := sb.String()
					if got != tt.want {
						t.Errorf("WrapInParenthesis() got = %v, want %v", got, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	type args struct {
		sw  io.StringWriter
		str string
		e   Expression
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should prepend",
			args: args{
				sw:  &strings.Builder{},
				str: "test",
				e:   testExpression{V: "1"},
			},
			want: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Prepend(tt.args.sw, tt.args.str, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sb, ok := tt.args.sw.(*strings.Builder); ok {
					got := sb.String()
					if got != tt.want {
						t.Errorf("Append() got = %v, want %v", got, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}
