package types

import "testing"

func TestObject_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		obj     *Object
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.obj.Serialize(); (err != nil) != tt.wantErr {
				t.Errorf("Object.Serialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
