package main

import (
	"testing"
)

func Test_envParse(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := envParse(); got == nil {
				t.Errorf("envParse() = %v", got)
			}
		})
	}
}
