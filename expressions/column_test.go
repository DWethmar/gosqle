package expressions

import (
	"io"
	"strings"
	"testing"
)

func TestColumn_Write(t *testing.T) {
	type fields struct {
		From string
		Name string
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
			name:   "no from",
			fields: fields{Name: "name"},
			args: args{
				sb: &strings.Builder{},
			},
			want:    "name",
			wantErr: false,
		},
		{
			name:   "astrisk",
			fields: fields{Name: "*", From: "table"},
			args: args{
				sb: &strings.Builder{},
			},
			want:    "table.*",
			wantErr: false,
		},
		{
			name:   "with from",
			fields: fields{From: "from", Name: "name"},
			args: args{
				sb: &strings.Builder{},
			},
			want:    "from.name",
			wantErr: false,
		},
		{
			name:   "should return error if no name",
			fields: fields{From: "from"},
			args: args{
				sb: &strings.Builder{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Column{
				From: tt.fields.From,
				Name: tt.fields.Name,
			}
			if err := s.WriteTo(tt.args.sb); (err != nil) != tt.wantErr {
				t.Errorf("Column.Write() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.want != tt.args.sb.String() {
				t.Errorf("Column.Write() got = %v, want %v", tt.args.sb.String(), tt.want)
			}
		})
	}
}

func TestColumn_WriteTo(t *testing.T) {
	type fields struct {
		From string
		Name string
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Column{
				From: tt.fields.From,
				Name: tt.fields.Name,
			}
			if err := s.WriteTo(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Column.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
