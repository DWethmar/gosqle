package postgres

import (
	"reflect"
	"strings"
	"testing"
)

func TestArgument_Write(t *testing.T) {
	type fields struct {
		Index int
		Value interface{}
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
			name: "should write to string builder",
			fields: fields{
				Index: 1,
				Value: "test",
			},
			args: args{
				sb: &strings.Builder{},
			},
			want:    "$1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Argument{
				Index: tt.fields.Index,
				Value: tt.fields.Value,
			}

			if err := s.Write(tt.args.sb); (err != nil) != tt.wantErr {
				t.Errorf("Argument.Write() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if sb := tt.args.sb; sb != nil {
					if str := sb.String(); str != tt.want {
						t.Errorf("Argument.Write() got = %q, want %q", str, tt.want)
					}
				}
			}
		})
	}
}

func TestNewArguments(t *testing.T) {
	t.Run("should return new arguments", func(t *testing.T) {
		if a := NewArguments(); a == nil {
			t.Errorf("NewArguments() should not be nil")
		}
	})
}

func TestArguments_NewArgument(t *testing.T) {
	type fields struct {
		Index int
		Args  []interface{}
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Argument
	}{
		{
			name: "should return new argument",
			fields: fields{
				Index: 0,
				Args:  []interface{}{},
			},
			args: args{
				value: "test",
			},
			want: &Argument{
				Index: 1,
				Value: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Arguments{
				Index:  tt.fields.Index,
				Values: tt.fields.Args,
			}
			if got := a.NewArgument(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Arguments.NewArgument() = %v, want %v", got, tt.want)
			}
		})
	}
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
			if got := NewArgument(tt.args.value, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArgumentAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("should increment index", func(t *testing.T) {
		args := NewArguments()

		for i := 0; i < 10; i++ {
			if got := args.NewArgument("test"); got.Index != i+1 {
				t.Errorf("NewArgumentAtIndex() = %v, want %v", got.Index, i+1)
			}
		}
	})
}
