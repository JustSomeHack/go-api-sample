package main

import (
	"testing"
)

func Test_printVersion(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Should print version of application",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printVersion()
		})
	}
}
