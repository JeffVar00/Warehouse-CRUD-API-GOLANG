package main

import (
	"context"
	"proyecto_inventarios/database"
	"proyecto_inventarios/internal/repository"
	"proyecto_inventarios/internal/service"
	"proyecto_inventarios/settings"

	"go.uber.org/fx"
)

//libreria de inyeccion de dependencias
//go get go.uber.org/fx

func main() {
	app := fx.New(
		fx.Provide( //le pasamos todas las funciones que devuelvan un stroke (error)
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(
		//TESTS PARA VER SI FUNCIONA
		//pueden exitir falsos positivos, ahora a probarlo con la bd
		// func(ctx context.Context, serv service.Service) {
		// 	err := serv.RegisterUser(ctx, "my@email.com", "my name", "my password")
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	u, err := serv.LoginUser(ctx, "my@email.com", "my password")
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	if u.Name != "my name" {
		// 		panic("wrong name")
		// 	}}
		),
	)
	app.Run()
}

//clean architecture concepto de forma de programar y organizarlos
//domain driven design dividir el proyecto en diferentes capas

//repository -> service -> api
//circular dependency, repository no puede acceder a service

//carpetas
//internal, cosas muy sensibles
//settings donde viene la configuracion de nuestro programa
//siempre es bueno tener dentro del package un archivo del mismo nombre y este es el punto de entrada
