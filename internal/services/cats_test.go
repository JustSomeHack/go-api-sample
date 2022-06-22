package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/one-byte-data/go-api-sample/internal/models"
	"gorm.io/gorm"
)

func Test_catsService_Add(t *testing.T) {
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
		// TODO: Add test cases.
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
