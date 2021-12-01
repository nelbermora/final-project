package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nelbermora/go-interfaces/internal/server"
)

func main() {
	// el router hace honor a su nombre y contiene los paths y rutas definidas,
	// y las respectivas funciones asociadas a cada ruta
	// el router es un objeto del tipo http.Handler, es transparente con cual libreria se genero (chi, mux, gin, native.. etc).
	router, _ := server.SetupDependencies()
	// se levanta el servidor indicando con el puerto y el router definido previamente
	log.Println("inicializando app")
	err := http.ListenAndServe(":18080", router)
	if err != nil {
		log.Fatal(fmt.Sprintf("could not initialize server: %s", err.Error()))
	}
}
