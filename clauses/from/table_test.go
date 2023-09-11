package from

import (
	"io"
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Table{
				Name: tt.fields.Name,
				As:   tt.fields.As,
			}
			if err := tr.WriteTo(tt.args.sw); (err != nil) != tt.wantErr {
				t.Errorf("Table.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
