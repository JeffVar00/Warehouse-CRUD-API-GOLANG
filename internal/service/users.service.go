package service

import (
	"context"
	"errors"
	"proyecto_inventarios/encryption"
	"proyecto_inventarios/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists") // no poner en mayuscula ni con punto
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRoleAlreadyExists  = errors.New("role already exists")
	ErrRoleDoesntExist    = errors.New("role doesnt exist")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {

	//primero checaremos si existe

	u, _ := s.repo.GetUserByEmail(ctx, email)

	if u != nil {
		return ErrUserAlreadyExists
	}

	//has password, usando la funcion de encryption
	bb, err := encryption.Encrypt([]byte(password))

	if err != nil {
		return err
	}

	//ecripto la contrasena
	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPass, err := encryption.Decrypt(bb)

	if err != nil {
		return nil, err
	}

	//hay que pasarla a string porque vieende ser un arreglo de bytes
	if string(decryptedPass) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) SaveUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)

	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleID == roleID {
			//retornar nill y hacer como que nada paso
			//return nil
			//o un mensaje de error
			return ErrRoleAlreadyExists
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)

	if err != nil {
		return err
	}

	role_found_flag := false

	for _, r := range roles {
		if r.RoleID == roleID {
			role_found_flag = true
			break
		}
	}

	if !role_found_flag {
		return ErrRoleDoesntExist
	}

	return s.repo.RemoveUserRole(ctx, userID, roleID)
}
