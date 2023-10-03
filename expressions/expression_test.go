package expressions

import (
	"io"
	"strings"
	"testing"
)

type testExpression struct {
	V string
}

func (e testExpression) WriteTo(writer io.StringWriter) error {
	_, err := writer.WriteString(e.V)
	return err
}

func TestList_WriteTo(t *testing.T) {
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
			if err := tt.e.WriteTo(tt.args.writer); (err != nil) != tt.wantErr {
				t.Errorf("List.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sb, ok := tt.args.writer.(*strings.Builder); ok {
					got := sb.String()
					if got != tt.want {
						t.Errorf("List.WriteTo() got = %v, want %v", got, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}
