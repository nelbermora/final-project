package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// probar el controller implica mockear un service
// creo un struct cualquiera
type mockSvc struct{}

// le asigno al struct creado el o los metodos que cumplan la interfaz del Service Real
func (m *mockSvc) DummyFunc() (string, error) {
	return "el que lo lea tiene un premio", nil
}

func TestEndpointController(t *testing.T) {

	svcMocked := &mockSvc{}
	myController := NewController(svcMocked)

	// para probar el controller necesito crear un request del tipo post, el cual
	// sera enviado al controller y comprobaremos que lo maneje correctamente
	// probaremos un request con no cumpla la validacion
	stringRequest := `{"input":"hol"}`
	req, _ := http.NewRequest("POST", "/endpoint", strings.NewReader(stringRequest))
	// el controller requiere de un request y un recorder para ser invocado
	// ya tenemos el request, entonces ahora creamos el recorder
	// para esto usamos la libreria httptest
	rr := httptest.NewRecorder()
	myController.EndpointController(rr, req)
	// segun nuestro desarrollo si la palabra input tiene menos de 4 letras
	// deberia retornar un error 400, procedemos a comprobarlo
	statusErrrExpected := 400
	assert.Equal(t, statusErrrExpected, rr.Result().StatusCode)

}
