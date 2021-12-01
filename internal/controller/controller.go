package controller

import (
	"encoding/json"
	"net/http"

	"github.com/nelbermora/go-interfaces/internal/model"
	"github.com/nelbermora/go-interfaces/internal/service"
)

// notar que en este caso no se hace uso de interfaz,
// esto debido a que los controladores tienen practicamente la misma definicion siemre
// son funcionoes que siempre recibe (http.ResponseWriter, r *http.Request)
type Controller struct {
	service service.Service
}

func NewController(svc service.Service) Controller {
	return Controller{
		service: svc,
	}
}

func (c *Controller) EndpointController(rw http.ResponseWriter, r *http.Request) {
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
	message, err := c.service.DummyFunc()
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
