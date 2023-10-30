package mysql

import (
	"reflect"
	"strings"
	"testing"
)

func TestArgument_Value(t *testing.T) {
	type fields struct {
		V interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "should return string value",
			fields: fields{
				V: "test",
			},
			want: "test",
		},
		{
			name: "should return int value",
			fields: fields{
				V: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Argument{
				V: tt.fields.V,
			}
			if got := s.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Argument.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgument_Write(t *testing.T) {
	type fields struct {
		V interface{}
	}
	type args struct {
		sb *strings.Builder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			name:   "should write string",
			fields: fields{V: "test"},
			args: args{
				sb: &strings.Builder{},
			},
			want:    "?",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Argument{
				V: tt.fields.V,
			}
			if err := s.Write(tt.args.sb); (err != nil) != tt.wantErr {
				t.Errorf("Argument.Write() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.want != tt.args.sb.String() {
				t.Errorf("Argument.Write() got = %v, want %v", tt.args.sb.String(), tt.want)
			}
		})
	}
}

func TestNewArgument(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *Argument
	}{
		{
			name: "should create new Argument",
			args: args{
				value: "test",
			},
			want: &Argument{
				V: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArgument(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArgument() = %v, want %v", got, tt.want)
			}
		})
	}
}
