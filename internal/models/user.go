package models

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

//como no voy a retornar el passsword no lo incluyo, porque esta es nuestra estructura de usuario
