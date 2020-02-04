package api

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test run success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run()
		})
	}
}
