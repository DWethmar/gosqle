package gosqle

import (
	"reflect"
	"strings"
	"testing"
)

func TestUpdate_ToSQL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		update  *Update
		want    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := new(strings.Builder)
			err := tt.update.Write(sb)

			if (err != nil) != tt.wantErr {
				t.Errorf("Update.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if query := sb.String(); query != tt.want {
				t.Errorf("Update.Write() query = %q, wantQuery %q", query, tt.want)
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
		want *Select
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
