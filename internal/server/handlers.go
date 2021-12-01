package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nelbermora/go-interfaces/internal/clients/db"
	"github.com/nelbermora/go-interfaces/internal/model"
	"github.com/nelbermora/go-interfaces/internal/repository"
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
		r.Post("/", endpointController)
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
	//se parsea el input del request sobre la variable req
	var req model.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	// si no se puede parseer entonces el json del request esta errado y se puede informar al cliente el bad request
	if err != nil {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(400)
		return
	}
	// si llegamos a este punto es porque el json de entrada esta bien
	// aca se puede hacer una validacion de calidad de dato si queremos
	err = req.Validate()
	if err != nil {
		response.Message = err.Error()
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(400)
		data, _ := json.Marshal(response)
		// se escribe la respuesta, este es el ultimo paso
		rw.Write(data)
		return
	}
	// si llegamos a este punto es porque todas las validaciones de la capa web estan ok
	// ahora para llamar al servicio debe crearse o instanciarse a traves de su funcion constructora
	//message, err := service.DummyFunc()
	// pero para crear un service debemos inyectarle el repo, y para el repo debemos inyectarle la DB que ya fue inicializada, no es una lindura?
	repo := repository.NewRepository(db.MyDB)
	service := service.NewService(repo)
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
