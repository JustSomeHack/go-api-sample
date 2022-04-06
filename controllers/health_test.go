package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestHealthGet(t *testing.T) {
	db, err := gorm.Open(postgres.Open("postgresql://root@cockroachdb:26257/defaultdb?sslmode=disable"))
	if err != nil {
		panic(fmt.Sprintf("unable to connect to database: %v", err))
	}

	router, err := SetupRouter(db)
	if err != nil {
		panic(err)
	}

	type args struct {
		method   string
		endpoint string
		body     interface{}
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantCode     int
	}{
		{
			name:         "Should pass health check",
			args:         args{method: "GET", endpoint: "/health", body: nil},
			wantResponse: "{\"message\":\"ok\"}",
			wantCode:     http.StatusOK,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.args.method, tt.args.endpoint, nil)
		router.ServeHTTP(w, req)

		if tt.wantCode != w.Code {
			t.Errorf("HealthGet() error = %v, wantCode %v", w.Code, tt.wantCode)
			return
		}

		if !reflect.DeepEqual(tt.wantResponse, w.Body.String()) {
			t.Errorf("HealthGet() error = %v, wantCode %v", w.Body.String(), tt.wantResponse)
		}
	}
}
