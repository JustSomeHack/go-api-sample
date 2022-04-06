package main

import (
	"os"
	"testing"
)

func Test_getConnectionString(t *testing.T) {
	tests := []struct {
		name      string
		param     string
		want      string
		wantPanic bool
	}{
		{
			name:      "Should get a valid connection string",
			param:     "postgresql://root@localhost:26257/defaultdb?sslmode=disable",
			want:      "postgresql://root@localhost:26257/defaultdb?sslmode=disable",
			wantPanic: false,
		},
		{
			name:      "Should panic",
			param:     "",
			want:      "",
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("CONNECTION_STRING", tt.param)

			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("getConnectionString() recover = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()

			if got := getConnectionString(); got != tt.want {
				t.Errorf("getConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_setupDatabase(t *testing.T) {
	type args struct {
		connectionString string
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "Should be able to connect to database",
			args: args{
				connectionString: "postgresql://root@cockroachdb:26257/defaultdb?sslmode=disable",
			},
			wantPanic: false,
		},
		{
			name: "Should panic",
			args: args{
				connectionString: "",
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("setupDatabase() recover = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()

			setupDatabase(tt.args.connectionString)
		})
	}
}
