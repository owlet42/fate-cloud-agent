package service

import (
	"fmt"
	"k8s.io/client-go/rest"
	"testing"
)

func TestGetPod(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPod()
		})
	}
}

func TestAbc(t *testing.T){
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.String())
}