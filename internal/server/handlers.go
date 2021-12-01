package server

import (
	"github.com/go-chi/chi"
	"github.com/nelbermora/go-interfaces/internal/clients/db"
	"github.com/nelbermora/go-interfaces/internal/controller"
	"github.com/nelbermora/go-interfaces/internal/repository"
	"github.com/nelbermora/go-interfaces/internal/service"
)

func setupHandlers(router chi.Router) {
	// recibo por parametros el router y defino que rutas quiero levantar
	// para este ejemplo se define una ruta llamada endpoint
	// se define el metodo get para dicha ruta
	// pueden definirse todos los metodos que apliquen
	// se genera el controlador, pero este requiere que se inyecte un service
	// pero para crear un service debemos inyectarle el repo, y para el repo debemos inyectarle la DB que ya fue inicializada, no es una lindura?
	repo := repository.NewRepository(db.MyDB)
	service := service.NewService(repo)
	myController := controller.NewController(service)
	router.Route("/endpoint", func(r chi.Router) {
		// cada endpoint y metodo debe estar vinculada a una funcion que controla lo que ocurre cuando un cliente
		// en este caso para el metodo GET se esta vinculando a la funcion controller o controladora llamada endpointController
		r.Post("/", myController.EndpointController)
	})
}
