package postgres

import (
	"reflect"
	"strings"
	"testing"
)

func TestArgument_Write(t *testing.T) {
	t.Run("should write argument", func(t *testing.T) {
		sb := new(strings.Builder)
		a := NewArgument(1, "test")
		if err := a.Write(sb); err != nil {
			t.Errorf("Argument.Write() error = %v", err)
		}

		if sb.String() != "$1" {
			t.Errorf("Argument.Write() got = %v, want %v", sb.String(), "$1")
		}
	})
}

func TestNewArguments(t *testing.T) {
	t.Run("should return new arguments", func(t *testing.T) {
		if a := NewArguments(); a == nil {
			t.Errorf("NewArguments() should not be nil")
		}
	})
}

func TestArguments_Create(t *testing.T) {
	t.Run("should create new argument", func(t *testing.T) {
		args := NewArguments()
		if a := args.Create("test"); a == nil {
			t.Errorf("NewArguments() should not be nil")
		}
	})
}

func TestNewArgument(t *testing.T) {
	type args struct {
		value interface{}
		index int
	}
	tests := []struct {
		name string
		args args
		want *Argument
	}{
		{
			name: "should return new argument",
			args: args{
				value: "test",
				index: 1,
			},
			want: &Argument{
				Index: 1,
				Value: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArgument(tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArgumentAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("should increment index", func(t *testing.T) {
		args := NewArguments()

		for i := 0; i < 10; i++ {
			a := args.Create(i)
			if err := a.Write(new(strings.Builder)); err != nil {
				t.Errorf("NewArgumentAtIndex() = %v", err)
			}
		}

		values := args.Values()

		if len(values) != 10 {
			t.Errorf("NewArgumentAtIndex() = %v, want %v", len(values), 10)
		}

		for i, v := range values {
			if v != i {
				t.Errorf("NewArgumentAtIndex() = %v, want %v", v, i)
			}
		}
	})
}
