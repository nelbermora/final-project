package service

import (
	"github.com/nelbermora/go-interfaces/internal/clients/db"
	"github.com/nelbermora/go-interfaces/internal/repository"
)

// en este pkg puede ir la logica de negocio

func DummyFunc() (string, error) {
	// para este ejemplo la funcion solo ejecutara la llamada al SP,
	// si esta se ejecuta correctamente entonces retornara un mensajem si no el error
	// ahora el repositorio se implementa a través de una interfaz, para invocar al repo debo crearlo con las dependencias que se requieran
	//return repository.LlamarSP()
	// esto debe hacerse al inicializarse la aplicacion, hacerlo aqui es una muy mala practica, pero a efectos del ejemplo lo haremos en este bloque de codigo
	// creo el repo indicando cual DB va a utilizar, le paso por parametro la base ya inicializada
	myRepo := repository.NewRepository(db.MyDB)
	return myRepo.LlamarSP()

}
