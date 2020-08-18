package config

// import (
// 	"reflect"
// 	"testing"
// )

// // func TestGet(t *testing.T) {
// // 	tests := []struct {
// // 		name string
// // 		want *Config
// // 	}{
// // 		{"xx", &Config{}},
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			if got := Get(); !reflect.DeepEqual(got, tt.want) {
// // 				t.Errorf("Get() = %v, want %v", got, tt.want)
// // 			}
// // 		})
// // 	}
// // }

// // func TestInit(t *testing.T) {
// // 	tests := []struct {
// // 		name string
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			Init()
// // 		})
// // 	}
// // }

// // func TestLoad(t *testing.T) {
// // 	type args struct {
// // 		runValidation bool
// // 	}
// // 	tests := []struct {
// // 		name    string
// // 		args    args
// // 		wantErr bool
// // 	}{
// // 		// TODO: Add test cases.
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			if err := Load(tt.args.runValidation); (err != nil) != tt.wantErr {
// // 				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
// // 			}
// // 		})
// // 	}
// // }

// func TestConfig_GetMatchingPathRule(t *testing.T) {
// 	type args struct {
// 		p string
// 	}
// 	tests := []struct {
// 		name    string
// 		c       Config
// 		args    args
// 		want    Rule
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.c.GetMatchingPathRule(tt.args.p)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Config.GetMatchingPathRule() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Config.GetMatchingPathRule() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestSetDefaultValues(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			SetDefaultValues()
// 		})
// 	}
// }

// func TestConfig_ValidateRules(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		c       Config
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.c.ValidateRules(); (err != nil) != tt.wantErr {
// 				t.Errorf("Config.ValidateRules() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
