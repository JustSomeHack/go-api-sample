package services

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/one-byte-data/go-api-sample/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_catsService_Add(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.Dialector{
		Config: &postgres.Config{Conn: db},
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	zeroID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	birth, _ := time.Parse(time.RFC3339, "2020-02-10T00:00:00Z")

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
			name:   "Add cat to database",
			fields: fields{db: gdb},
			args: args{ctx: context.Background(), cat: &models.Cat{
				Name:      "Nacho",
				Breed:     "Tabby",
				Color:     "Orange",
				Birthdate: birth,
				Weight:    17,
			}},
			want:    &zeroID,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(`INSERT INTO "cats"`).WithArgs(
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.Dialector{
		Config: &postgres.Config{Conn: db},
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

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
			name:    "Delete cat from database",
			fields:  fields{db: gdb},
			args:    args{ctx: context.Background(), id: uuid.New()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(`DELETE FROM "cats" WHERE`).WithArgs(tt.args.id).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.Dialector{
		Config: &postgres.Config{Conn: db},
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		filter interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Cat
		wantErr bool
	}{
		{
			name:    "Get all cats",
			fields:  fields{db: gdb},
			args:    args{ctx: context.Background(), filter: nil},
			want:    []models.Cat{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery(`SELECT`).WillReturnRows(sqlmock.NewRows([]string{""}))
			s := &catsService{
				db: tt.fields.db,
			}
			got, err := s.Get(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("catsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("catsService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
