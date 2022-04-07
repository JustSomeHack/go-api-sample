package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/JustSomeHack/go-api-sample/internal/models"
	"github.com/JustSomeHack/go-api-sample/cmd/tests"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
)

func TestDogsDelete(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	router, err := SetupRouter(tests.DB)
	if err != nil {
		panic(err)
	}

	type args struct {
		method   string
		endpoint string
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
			},
			wantResponse: fmt.Sprintf("{\"deleted\":\"%s\"}", tests.Dogs[0].ID.String()),
			wantCode:     http.StatusOK,
		},
		{
			name: "Should not delete an invalid ID",
			args: args{
				method:   "DELETE",
				endpoint: fmt.Sprintf("/dogs/%s", "not_a_valid_id"),
			},
			wantResponse: "{\"message\":\"invalid id\"}",
			wantCode:     http.StatusBadRequest,
		},
		{
			name: "Should not delete an ID that does not exist",
			args: args{
				method:   "DELETE",
				endpoint: fmt.Sprintf("/dogs/%s", uuid.New().String()),
			},
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

func TestDogsGet(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	router, err := SetupRouter(tests.DB)
	if err != nil {
		panic(err)
	}

	type args struct {
		method   string
		endpoint string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantCode  int
	}{
		{
			name: "Should all dogs",
			args: args{
				method:   "GET",
				endpoint: "/dogs",
			},
			wantCount: len(tests.Dogs),
			wantCode:  http.StatusOK,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.args.method, tt.args.endpoint, nil)
		router.ServeHTTP(w, req)

		if tt.wantCode != w.Code {
			t.Errorf("DogsGet() error = %v, wantCode %v", w.Code, tt.wantCode)
			return
		}

		dogs := make([]models.Dog, 0)
		err := json.Unmarshal(w.Body.Bytes(), &dogs)
		if err != nil {
			t.Errorf("DogsGet() error = %v, wantCount %v", err, tt.wantCount)
			return
		}

		if tt.wantCount != len(dogs) {
			t.Errorf("DogsGet() error = %v, wantCount %v", len(dogs), tt.wantCount)
		}
	}
}

func TestDogsGetOne(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	router, err := SetupRouter(tests.DB)
	if err != nil {
		panic(err)
	}

	type args struct {
		method   string
		endpoint string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse *models.Dog
		wantCode     int
	}{
		{
			name: "Should get a dog by ID",
			args: args{
				method:   "GET",
				endpoint: fmt.Sprintf("/dogs/%s", tests.Dogs[0].ID.String()),
			},
			wantResponse: &tests.Dogs[0],
			wantCode:     http.StatusOK,
		},
		{
			name: "Should not get an invalid ID",
			args: args{
				method:   "GET",
				endpoint: fmt.Sprintf("/dogs/%s", "not_a_valid_id"),
			},
			wantResponse: nil,
			wantCode:     http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.args.method, tt.args.endpoint, nil)
		router.ServeHTTP(w, req)

		if tt.wantCode != w.Code {
			t.Errorf("DogsGetOne() error = %v, wantCode %v", w.Code, tt.wantCode)
			return
		}

		if tt.wantResponse != nil {
			dog := new(models.Dog)
			err := json.Unmarshal(w.Body.Bytes(), &dog)
			if err != nil {
				t.Errorf("DogsGetOne() error = %v, wantCount %v", err, tt.wantResponse)
				return
			}

			if !reflect.DeepEqual(tt.wantResponse.ID, dog.ID) {
				t.Errorf("DogsGetOne() error = %v, wantCode %v", dog, tt.wantResponse)
			}
		}
	}
}

func TestDogsPost(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	router, err := SetupRouter(tests.DB)
	if err != nil {
		panic(err)
	}

	type args struct {
		method   string
		endpoint string
		body     *models.Dog
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantCode     int
	}{
		{
			name: "Should add a dog to the database",
			args: args{
				method:   "POST",
				endpoint: "/dogs",
				body: &models.Dog{
					Name:      "Spike",
					Breed:     "Bulldog",
					Color:     "Grey",
					Birthdate: time.Date(2021, 12, 10, 0, 0, 0, 0, time.UTC),
					Weight:    55,
				}},
			wantResponse: "created",
			wantCode:     http.StatusCreated,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.args.body)
		if err != nil {
			t.Errorf("DogsPost() unable to marshal %v", tt.args.body)
			return
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.args.method, tt.args.endpoint, bytes.NewReader(data))
		req.Header.Add("Content-type", "application/json")
		router.ServeHTTP(w, req)

		if tt.wantCode != w.Code {
			t.Errorf("DogsPost() error = %v, wantCode %v", w.Code, tt.wantCode)
			return
		}

		if !strings.Contains(w.Body.String(), tt.wantResponse) {
			t.Errorf("DogsPost() error = %v, wantCode %v", w.Body.String(), tt.wantResponse)
		}
	}
}

func TestDogsPut(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	router, err := SetupRouter(tests.DB)
	if err != nil {
		panic(err)
	}

	type args struct {
		method   string
		endpoint string
		body     *models.Dog
	}
	tests := []struct {
		name         string
		args         args
		wantResponse string
		wantCode     int
	}{
		{
			name: "Should update a dog",
			args: args{
				method:   "PUT",
				endpoint: fmt.Sprintf("/dogs/%s", tests.Dogs[0].ID.String()),
				body: &models.Dog{
					Name:      "Nacho",
					Breed:     "Tabby",
					Color:     "Orange",
					Birthdate: time.Date(2020, 2, 10, 0, 0, 0, 0, time.UTC),
					Weight:    20,
				}},
			wantResponse: "updated",
			wantCode:     http.StatusAccepted,
		},
		{
			name: "Should not update an invalid id",
			args: args{
				method:   "PUT",
				endpoint: fmt.Sprintf("/dogs/%s", "invalid_id_here"),
				body: &models.Dog{
					Name:      "Nacho",
					Breed:     "Tabby",
					Color:     "Orange",
					Birthdate: time.Date(2020, 2, 10, 0, 0, 0, 0, time.UTC),
					Weight:    20,
				}},
			wantResponse: "invalid",
			wantCode:     http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.args.body)
		if err != nil {
			t.Errorf("DogsPut() unable to marshal %v", tt.args.body)
			return
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(tt.args.method, tt.args.endpoint, bytes.NewReader(data))
		req.Header.Add("Content-type", "application/json")
		router.ServeHTTP(w, req)

		if tt.wantCode != w.Code {
			t.Errorf("DogsPut() error = %v, wantCode %v", w.Code, tt.wantCode)
			return
		}

		if !strings.Contains(w.Body.String(), tt.wantResponse) {
			t.Errorf("DogsPut() error = %v, wantCode %v", w.Body.String(), tt.wantResponse)
		}
	}
}
