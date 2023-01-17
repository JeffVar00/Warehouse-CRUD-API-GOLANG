package repository

import (
	"context"
	"proyecto_inventarios/internal/entity"
)

const (
	qryInsertUser = `INSERT INTO USERS (email, name, password) VALUES (?,?,?)`

	qryGetUserByEmail = `SELECT id, email, name, password FROM USERS WHERE email = ?`

	//no es necesario colocar :user_id pero como lo colocamos en el entity nos ayuda a identificar, puede ser colocado el ?
	qryInsertUserRole = `INSERT INTO USER_ROLES (user_id, role_id) VALUES (:user_id, :role_id)`

	qryDeleteUserRole = `DELETE FROM USER_ROLES WHERE user_id = ? AND role_id = ?`
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

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)
	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	_, err := r.db.NamedExecContext(ctx, qryDeleteUserRole, data)
	return err
}

func (r *repo) GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error) {

	roles := []entity.UserRole{}
	//funcion que viene de SQL
	err := r.db.SelectContext(ctx, &roles, "select user_id, role_id from USER_ROLES where user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
