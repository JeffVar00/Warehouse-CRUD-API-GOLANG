package database

import (
	"context"
	"fmt"
	"proyecto_inventarios/settings"

	_ "github.com/go-sql-driver/mysql" //el _ lo inicialia de forma automatica y activa un init si tiene alguno
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name)

	return sqlx.ConnectContext(ctx, "mysql", connectionString)
	//le estamos mandando un driver pero no lo tenemos instalado
	//comando:
	//go get github.com/go-sql-driver/mysql
}
