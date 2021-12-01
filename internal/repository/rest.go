package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nelbermora/go-interfaces/internal/model"
)

func LlamarRest() (string, error) {
	// se general el struct u objeto del request
	request := generateRequest()
	// se genera el json de ese objeto
	json_data, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	// se ejecuta el request
	doReq, err := http.Post("dummyUrl", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		return "", err
	}
	// por definicion la libreria http no requiere hacer un close del body, pero somos unos paranoicos y por las dudas lo hacemos
	defer doReq.Body.Close()
	// se define variable que almacenará la respuesta
	var response model.ApiResponse
	// se parsea el json de la respuesta a la variable response
	json.NewDecoder(doReq.Body).Decode(&response)
	fmt.Println(response)
	if response.Codigo != 0 {
		return "", errors.New("error en API de precios")
	}
	return "Api invocado Correctamente", nil
}

func generateRequest() *model.ApiRequest {
	return &model.ApiRequest{
		Canal:    "CP",
		SubCanal: "00",
		Datos: model.DatosRequest{
			Currency:  "ARS",
			StartDate: "2020-10-01",
			EndDate:   "2025-10-01",
			Use:       "COMMISSION",
			Assets:    []string{"80436", "0019"},
		},
	}
}
