package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/JustSomeHack/go-api-sample/tests"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
)

func TestDogsDelete(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(connectionString))
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
			name: "Should delete a dog by ID",
			args: args{
				method:   "DELETE",
				endpoint: fmt.Sprintf("/dogs/%s", tests.Dogs[0].ID.String()),
				body:     nil},
			wantResponse: fmt.Sprintf("{\"deleted\":\"%s\"}", tests.Dogs[0].ID.String()),
			wantCode:     http.StatusOK,
		},
		{
			name: "Should not delete an invalid ID",
			args: args{
				method:   "DELETE",
				endpoint: fmt.Sprintf("/dogs/%s", "not_a_valid_id"),
				body:     nil},
			wantResponse: "{\"message\":\"invalid id\"}",
			wantCode:     http.StatusBadRequest,
		},
		{
			name: "Should not delete an ID that does not exist",
			args: args{
				method:   "DELETE",
				endpoint: fmt.Sprintf("/dogs/%s", uuid.New().String()),
				body:     nil},
			wantResponse: "{\"message\":\"there was an error\"}",
			wantCode:     http.StatusInternalServerError,
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
