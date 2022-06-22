package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/one-byte-data/go-api-sample/internal/models"
	"gorm.io/gorm"
)

func Test_dogsService_Add(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		dog *models.Dog
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &dogsService{
				db: tt.fields.db,
			}
			got, err := s.Add(tt.args.ctx, tt.args.dog)
			if (err != nil) != tt.wantErr {
				t.Errorf("dogsService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dogsService.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
