package repository

import (
	"context"
	"proyecto_inventarios/internal/entity"
)

const (
	qryInsertUser = `INSERT INTO USERS (email, name, password) VALUES (?,?,?)`

	qryGetUserByEmail = `SELECT id, email, name, password FROM USERS WHERE email = ?`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err //siempre cifrar contrasenas
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	u := &entity.User{}

	//clear better than clever

	//get context pide una referencia de una variable entonces es importante que lo reciba como puntero
	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)

	if err != nil {
		return nil, err
	}

	return u, nil
}
