package predicates

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
)

func TestComparison_Write(t *testing.T) {
	type fields struct {
		Left     expressions.Expression
		Operator string
		Right    expressions.Expression
	}
	type args struct {
		writer io.StringWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should left and right",
			fields: fields{
				Left:     expressions.String("left"),
				Operator: "=",
				Right:    expressions.String("right"),
			},
			args: args{
				writer: new(strings.Builder),
			},
			want:    "left = right",
			wantErr: false,
		},
		{
			name: "should error when writer is nil",
			fields: fields{
				Left:     expressions.String("left"),
				Operator: "=",
				Right:    expressions.String("right"),
			},
			args: args{
				writer: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should error when left is nil",
			fields: fields{
				Left:     nil,
				Operator: "=",
				Right:    expressions.String("right"),
			},
			args: args{
				writer: new(strings.Builder),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should error when operator is empty",
			fields: fields{
				Left:     expressions.String("left"),
				Operator: "",
				Right:    expressions.String("right"),
			},
			args: args{
				writer: new(strings.Builder),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should error when right is nil",
			fields: fields{
				Left:     expressions.String("left"),
				Operator: "=",
				Right:    nil,
			},
			args: args{
				writer: new(strings.Builder),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Comparison{
				Left:     tt.fields.Left,
				Operator: tt.fields.Operator,
				Right:    tt.fields.Right,
			}
			if err := c.Write(tt.args.writer); (err != nil) != tt.wantErr {
				t.Errorf("Comparison.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != "" {
				if sb, ok := tt.args.writer.(*strings.Builder); ok {
					if s := sb.String(); s != tt.want {
						t.Errorf("Comparison.Write() = %v, want %v", s, tt.want)
					}
				} else {
					t.Errorf("expected a strings.Builder")
				}
			}
		})
	}
}

func TestEQ(t *testing.T) {
	type args struct {
		left  expressions.Expression
		right expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should create comparison",
			args: args{
				left:  expressions.String("left"),
				right: expressions.String("right"),
			},
			want: &Comparison{
				Left:     expressions.String("left"),
				Operator: "=",
				Right:    expressions.String("right"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EQ(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNE(t *testing.T) {
	type args struct {
		left  expressions.Expression
		right expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should create comparison",
			args: args{
				left:  expressions.String("left"),
				right: expressions.String("right"),
			},
			want: &Comparison{
				Left:     expressions.String("left"),
				Operator: "!=",
				Right:    expressions.String("right"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NE(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGT(t *testing.T) {
	type args struct {
		left  expressions.Expression
		right expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should create comparison",
			args: args{
				left:  expressions.String("left"),
				right: expressions.String("right"),
			},
			want: &Comparison{
				Left:     expressions.String("left"),
				Operator: ">",
				Right:    expressions.String("right"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GT(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGTE(t *testing.T) {
	type args struct {
		left  expressions.Expression
		right expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should create comparison",
			args: args{
				left:  expressions.String("left"),
				right: expressions.String("right"),
			},
			want: &Comparison{
				Left:     expressions.String("left"),
				Operator: ">=",
				Right:    expressions.String("right"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GTE(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLT(t *testing.T) {
	type args struct {
		left  expressions.Expression
		right expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should create comparison",
			args: args{
				left:  expressions.String("left"),
				right: expressions.String("right"),
			},
			want: &Comparison{
				Left:     expressions.String("left"),
				Operator: "<",
				Right:    expressions.String("right"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LT(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLTE(t *testing.T) {
	type args struct {
		left  expressions.Expression
		right expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should create comparison",
			args: args{
				left:  expressions.String("left"),
				right: expressions.String("right"),
			},
			want: &Comparison{
				Left:     expressions.String("left"),
				Operator: "<=",
				Right:    expressions.String("right"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LTE(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLike(t *testing.T) {
	type args struct {
		left  expressions.Expression
		right expressions.Expression
	}
	tests := []struct {
		name string
		args args
		want *Comparison
	}{
		{
			name: "should create comparison",
			args: args{
				left:  expressions.String("left"),
				right: expressions.String("right"),
			},
			want: &Comparison{
				Left:     expressions.String("left"),
				Operator: "LIKE",
				Right:    expressions.String("right"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Like(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Like() = %v, want %v", got, tt.want)
			}
		})
	}
}
