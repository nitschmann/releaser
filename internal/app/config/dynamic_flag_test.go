package config

import "testing"

func TestDynamicFlag_Validate(t *testing.T) {
	tests := []struct {
		name    string
		d       DynamicFlag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("DynamicFlag.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
