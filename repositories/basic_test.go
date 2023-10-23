package repositories

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/gsdenys/go-fiber-postgres/database"
	"github.com/gsdenys/go-fiber-postgres/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_validateEntity(t *testing.T) {
	type args struct {
		entity interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Entity",
			args: args{
				entity: struct {
					Name  string `validate:"required"`
					Email string `validate:"required,email"`
				}{
					Name:  "John Doe",
					Email: "john@example.com",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Entity",
			args: args{
				entity: struct {
					Name  string `validate:"required"`
					Email string `validate:"required,email"`
				}{
					Name:  "",              // Invalid: Name is required
					Email: "invalid-email", // Invalid: Email format is incorrect
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateEntity(tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("validateEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		{
			name: "New Repository Instance",
			args: args{
				db: &gorm.DB{}, // Mock Gorm DB instance
			},
			want: &Repository{
				db: &gorm.DB{}, // Expected Repository instance with the same DB instance
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CreateEntity(t *testing.T) {

	t.Setenv("DATABASE_DSN", "user=postgres dbname=postgres password=postgres host=localhost port=5432 sslmode=disable")

	err := database.InitDB()
	assert.Nil(t, err, "the database should be initialized")

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		entity interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create Valid Entity",
			fields: fields{
				db: database.DB,
			},
			args: args{
				entity: &models.User{
					Username: uuid.NewString(),
					Email:    fmt.Sprintf("%s@email.com", uuid.NewString()),
				},
			},
			wantErr: false,
		},
		{
			name: "Create Invalid Entity",
			fields: fields{
				db: database.DB,
			},
			args: args{
				entity: &models.User{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			if err := r.CreateEntity(tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetEntityByID(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		entity interface{}
		id     uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			if err := r.GetEntityByID(tt.args.entity, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetEntityByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
