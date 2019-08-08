package types

import "testing"

func TestExtention(t *testing.T) {
	type args struct {
		mimeType string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Extention(tt.args.mimeType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extention() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Extention() = %v, want %v", got, tt.want)
			}
		})
	}
}
