package statement

import (
	"io"
	"strings"
	"testing"
)

func TestWriteUpdate(t *testing.T) {
	t.Run("should write UPDATE", func(t *testing.T) {
		sb := new(strings.Builder)
		if err := WriteUpdate(sb, "table"); err != nil {
			t.Errorf("WriteUpdate() error = %v", err)
		}

		if sb.String() != "UPDATE table" {
			t.Errorf("WriteUpdate() got = %v, want %v", sb.String(), "UPDATE table")
		}
	})
}

func TestUpdate_Write(t *testing.T) {
	type fields struct {
		ClauseWriter ClauseWriter
		table        string
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
			name: "should write UPDATE",
			fields: fields{
				table:        "table",
				ClauseWriter: NewUpdateClauseWriter(),
			},
			args: args{
				sw: new(strings.Builder),
			},
			want:    "UPDATE table",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Update{
				ClauseWriter: tt.fields.ClauseWriter,
				table:        tt.fields.table,
			}

			if err := u.Write(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Update.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			if sb, ok := tt.args.sw.(*strings.Builder); ok {
				if got := sb.String(); got != tt.want {
					t.Errorf("Update.Write() = %v, want %v", got, tt.want)
				}
			} else if tt.want != "" {
				t.Errorf("expected string builder")
			}
		})
	}
}
