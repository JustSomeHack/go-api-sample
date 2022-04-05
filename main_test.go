package main

import "testing"

func Test_printMessage(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sample Test",
			args: args{message: "Test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printMessage(tt.args.message)
		})
	}
}
