package entity

type User struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

//necesito una estrucutra para el usuario entonces agregamos el tag de json
//- en el json significa que esta oculto
