package service

import (
	context "context"
	"os"
	"proyecto_inventarios/encryption"
	"proyecto_inventarios/internal/entity"
	"proyecto_inventarios/internal/repository"
	"testing"

	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

//cntrl shitf t + coverage para ver que se contempla en el testing, esto se ve en el user.service.go
//esto se concentra en a logica

func TestMain(m *testing.M) {
	//_ es para que ignore el error o la variable
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)

	u := &entity.User{Email: "test@exists.com", Password: encryptedPassword}

	//no es bueno setear los valores del testo dentro de los mismos tests es mejor de forma global
	repo = &repository.MockRepository{}
	//si le paso este correo esta interpretando un correo que no existe y hace de cuenta que no retorna nada
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	//en caso contrario aqui es un caso de que un usuario existe entonces retornar ese usuario
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(u, nil)
	//ahora con el primer caso de error al guardar un usuario y siempre e snull y recibe lo que sea
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run() //devuelve un codigo, el status code
	os.Exit(code)   //si algunio de los unit test devuelve algo diferente de 0 lo permite devolver en consola
}

//aqui hacemos los testeos con mockery
//UNIT TESTS
//VIDEO DE EXPLICACION EN EL CANAL DE GO SIMPLIFICADO
//PARA HACER LOS TESTS DENTRO DE LA CARPETA DE SERVICE HACER go test O VISUAL ESTUDIO PERMITE HACERLOS DIRECTAMENTE

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		UserName      string
		ExpectedError error
	}{
		{
			Name:          "Resgister_User_Success",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "Resgister_User_UserAlreadyExists",
			Email:         "test@exists.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	//* *entity es para que no se confunda con el entity del paquete, es un tipo de dato
	//* &entity es para que sea un puntero a ese tipo de datos y acceder a sus archivos

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}

}

//se esta creando una nueva referencia de cada uno de los indices para poder trabajar con unit test paralelo

//ahora

// ok  	proyecto_inventarios/internal/service	0.271s
// si funciono
func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "Login_User_Success",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: nil, //esto es po defecto pero bueno ponerlo para entender
		},
		{
			Name:          "Login_User_InvalidPassword",
			Email:         "test@exists.com",
			Password:      "invalidPassword",
			ExpectedError: ErrInvalidCredentials,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("expected error %v, got %v", tc.ExpectedError, err)
			}

		})
	}

}
