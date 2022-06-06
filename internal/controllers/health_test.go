package controllers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/one-byte-data/go-api-sample/cmd/tests"
	"gorm.io/driver/postgres"
)

func TestHealthGet(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	router, err := SetupRouter(tests.DB)
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
