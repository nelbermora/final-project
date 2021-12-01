package server

/*
 en el package server se define todo lo asociado a la capa web, es decir, del servidor.
 todo lo referente, a construccion de router, definicion de paths o endpoints,
 y vinculacion entre endpoints y controllers se hace en el package server
*/

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nelbermora/go-interfaces/config"
	"github.com/nelbermora/go-interfaces/internal/clients/db"
)

// esta func sirve para inicializar las dependencias necesarias
func SetupDependencies() (http.Handler, error) {
	// se utiliza chi para crear el router
	// en nuestro template oficial hay una funcion del template que se encarga de crear el router por nosotros
	cfg := config.ResolveConfiguration()
	router := chi.NewRouter()
	// se inicializa la db
	db.InitializeDB(cfg)
	// en la funcion setupHandlers se hace la definicion de cada endpoint y su funcion controladora asociada
	setupHandlers(router)

	// se inicializa la db
	//para este ejemplo puedo comentar la inicializacion de la db
	//db.InitializeDB()
	return router, nil
}
