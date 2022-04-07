package services

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/JustSomeHack/go-api-sample/internal/models"
	"github.com/JustSomeHack/go-api-sample/cmd/tests"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewCatsService(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want CatsService
	}{
		{
			name: "Should get valid interface back",
			args: args{db: tests.DB},
			want: &catsService{db: tests.DB},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCatsService(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCatsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_catsService_Add(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	catID := uuid.New()

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		cat *models.Cat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *uuid.UUID
		wantErr bool
	}{
		{
			name:   "Should add a cat to the database",
			fields: fields{db: tests.DB},
			args: args{
				ctx: context.Background(),
				cat: &models.Cat{
					ID:        catID,
					Name:      "Nacho",
					Breed:     "Tabby",
					Color:     "Orange",
					Birthdate: time.Now(),
					Weight:    17,
				},
			},
			want:    &catID,
			wantErr: false,
		},
		{
			name:   "Should not be able to add empty cat",
			fields: fields{db: tests.DB},
			args: args{
				ctx: context.Background(),
				cat: &models.Cat{
					ID: uuid.New(),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catsService{
				db: tt.fields.db,
			}
			got, err := s.Add(tt.args.ctx, tt.args.cat)
			if (err != nil) != tt.wantErr {
				t.Errorf("catsService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("catsService.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_catsService_Delete(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Should delete cat by ID",
			fields:  fields{db: tests.DB},
			args:    args{ctx: context.Background(), id: tests.Cats[0].ID},
			wantErr: false,
		},
		{
			name:    "Should not delete an ID that does not exists",
			fields:  fields{db: tests.DB},
			args:    args{ctx: context.Background(), id: uuid.New()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catsService{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("catsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_catsService_Get(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		filter interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "Should get all the cats",
			fields:    fields{db: tests.DB},
			args:      args{ctx: context.Background(), filter: nil},
			wantCount: 3,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catsService{
				db: tt.fields.db,
			}
			got, err := s.Get(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("catsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantCount {
				t.Errorf("catsService.Get() = %v, want %v", len(got), tt.wantCount)
			}
		})
	}
}

func Test_catsService_GetOne(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Cat
		wantErr bool
	}{
		{
			name:    "Should get cat by ID",
			fields:  fields{db: tests.DB},
			args:    args{ctx: context.Background(), id: tests.Cats[0].ID},
			want:    &tests.Cats[0],
			wantErr: false,
		},
		{
			name:    "Should not get cat that does not exist",
			fields:  fields{db: tests.DB},
			args:    args{ctx: context.Background(), id: uuid.New()},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catsService{
				db: tt.fields.db,
			}
			got, err := s.GetOne(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("catsService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got.ID, tt.want.ID) {
					t.Errorf("catsService.GetOne() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_catsService_Update(t *testing.T) {
	teardownTests := tests.SetupTests(t, postgres.Open(tests.ConnectionString))
	defer teardownTests(t)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
		cat *models.Cat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Should update a cat by ID",
			fields: fields{db: tests.DB},
			args: args{ctx: context.Background(), id: tests.Cats[0].ID, cat: &models.Cat{
				Name:      "Nacho",
				Breed:     "Tabby",
				Color:     "Orange",
				Birthdate: time.Date(2020, 2, 10, 0, 0, 0, 0, time.UTC),
				Weight:    19,
			}},
			wantErr: false,
		},
		{
			name:   "Should not update a cat with no valid ID",
			fields: fields{db: tests.DB},
			args: args{ctx: context.Background(), id: uuid.New(), cat: &models.Cat{
				Name:      "Nacho",
				Breed:     "Tabby",
				Color:     "Orange",
				Birthdate: time.Date(2020, 2, 10, 0, 0, 0, 0, time.UTC),
				Weight:    19,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catsService{
				db: tt.fields.db,
			}
			if err := s.Update(tt.args.ctx, tt.args.id, tt.args.cat); (err != nil) != tt.wantErr {
				t.Errorf("catsService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
