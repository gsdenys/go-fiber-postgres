package database

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	tests := []struct {
		name    string
		dsn     string
		wantErr bool
	}{
		{
			name:    "Valid DSN",
			dsn:     "user=postgres dbname=postgres password=postgres host=localhost port=5432 sslmode=disable",
			wantErr: false,
		},
		{
			name:    "Invalid DSN - Missing Password",
			dsn:     "user=postgres dbname=postgres host=localhost port=5432 sslmode=disable",
			wantErr: true,
		},
		{
			name:    "Invalid DSN - Empty DSN",
			dsn:     "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.dsn) > 0 {
				t.Setenv(dsnEnvVar, tt.dsn)
			}

			if err := InitDB(); (err != nil) != tt.wantErr {
				t.Errorf("InitDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
