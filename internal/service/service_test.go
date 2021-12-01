package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// aca debemos mockear el repo para probar el service.
// para esto generamos un objeto que responda la info que necesitamos
// solo hay que garantizar que cumpla con la interfaz
// creo un struct que hara de repo
type mockRepo struct{}

// defino los metodos identicos a los que debe tener el repo
func (mr *mockRepo) LlamarSP() (string, error) {
	return "repo que devuelve lo que quiero", nil
}
func TestDummyFunc(t *testing.T) {
	// para probar el repo debemos instanciarlo
	// lo hacemos inyectando nuestro repo mockeado
	repo := &mockRepo{}
	svc := NewService(repo)
	// ahora invocamos al service controlando que responda exactamente lo que el repo le informa
	stringObtenido, err := svc.DummyFunc()
	// validamos que las respuestas esten correctas segun lo esperado
	assert.Equal(t, "repo que devuelve lo que quiero", stringObtenido)
	assert.Nil(t, err)
}
