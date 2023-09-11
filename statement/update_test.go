package statement

import (
	"io"
	"reflect"
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

func TestUpdate_WriteTo(t *testing.T) {
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Update{
				ClauseWriter: tt.fields.ClauseWriter,
				table:        tt.fields.table,
			}
			if err := u.WriteTo(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Update.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewUpdate(t *testing.T) {
	type args struct {
		table string
	}
	tests := []struct {
		name string
		args args
		want *Update
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUpdate(tt.args.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
