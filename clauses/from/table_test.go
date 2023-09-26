package from

import (
	"io"
	"strings"
	"testing"
)

func TestTable_WriteTo(t *testing.T) {
	type fields struct {
		Name string
		As   string
	}
	type args struct {
		sw io.StringWriter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "write table",
			fields: fields{
				Name: "table",
			},
			args: args{
				sw: new(strings.Builder),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Table(tt.fields.Name)

			if err := tr.WriteTo(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("tr.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if tt.want != "" {
					if sb, ok := tt.args.sw.(*strings.Builder); ok {
						if got := sb.String(); got != tt.want {
							t.Errorf("tr.WriteTo() = %q, want %q", got, tt.want)
						}
					} else {
						t.Errorf("expected string builder")
					}
				}
			}
		})
	}
}
