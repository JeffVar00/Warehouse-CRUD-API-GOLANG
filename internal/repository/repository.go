package repository

import (
	"context"
	"proyecto_inventarios/internal/entity"

	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps the basic CRUD methods
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
type repo struct {
	db *sqlx.DB
}

// nunca deberiamos retornar interfaces, debemos recibir las interfaces y reportar structs
func New(db *sqlx.DB) Repository {
	return &repo{db: db}
}

//el error decia que no tenia implementados los metodos del repositori del mockery
//instalar mockery
//go install github.com/vektra/mockery/v2@latest
//go generate ./... //esto lo hace de forma recursiva
//y crea el mock.repository.go
