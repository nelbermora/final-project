package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nelbermora/go-interfaces/internal/model"
	"github.com/nelbermora/go-interfaces/internal/service"
)

func setupHandlers(router chi.Router) {
	// recibo por parametros el router y defino que rutas quiero levantar
	// para este ejemplo se define una ruta llamada endpoint
	// se define el metodo get para dicha ruta
	// pueden definirse todos los metodos que apliquen
	router.Route("/endpoint", func(r chi.Router) {
		// cada endpoint y metodo debe estar vinculada a una funcion que controla lo que ocurre cuando un cliente
		// en este caso para el metodo GET se esta vinculando a la funcion controller o controladora llamada endpointController
		r.Get("/", endpointController)
	})
}

func endpointController(rw http.ResponseWriter, r *http.Request) {
	/*
		en esta funcion se controla el request y response,
		aca se pueden hacer valdiaciones del request y si esta todo ok
		se invoca a la funcion que contiene la logica de negocio.
		Tambien se general las respuestas a cada peticion
		El alcance de este ejemplo es solo demostrar el funcionamiento de la capa web o http a traves de go chi
	*/
	// para efectos ilustrativos se define un modelo que sera nuestra respuesta
	// se crea el objeto que vamos a responder
	var response model.Response
	message, err := service.DummyFunc()
	if err != nil {
		response.Message = "Error al consultar DB: " + err.Error()
	} else {
		response.Message = message
	}
	// se añaden los headers correspondientes y el status de la respuesta
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(200)
	// se parsea a json el struct Response
	data, _ := json.Marshal(response)
	// se escribe la respuesta, este es el ultimo paso
	rw.Write(data)

}
