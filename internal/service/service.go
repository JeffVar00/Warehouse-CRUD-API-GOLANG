package service

import (
	"context"
	"proyecto_inventarios/internal/models"
	"proyecto_inventarios/internal/repository"
)

// Service is the bussiness logic of the application
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	//devuelvo un model porque ahora es al capa de services encesito retornar el modelo del usuario sin el password
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{repo: repo}
}
