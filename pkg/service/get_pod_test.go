package service

import (
	"testing"
)

func TestGetPod(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name:"a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPod()
		})
	}
}
