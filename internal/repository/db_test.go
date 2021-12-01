package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// vamos a probar el archivo repository
// se debe testear la mayor parte de codigo que se pueda
// pueden hacerse diferentes tests para una misma funcion
// ej: podemos probar el metodo llamarSP distintas veces con distintas condiciones
func TestLlamarSP(t *testing.T) {
	// para probar el repository es necesario 'instanciar' o crear un nuevo repo
	// para esto llamamos a la funcion constructura, pero antes necesitaremos una db mockeada

	// mockear cualquier capa es sencillo, pero para mockear una DB usaremos el pkg sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' no se pudo levantar el mock", err)
	}
	defer db.Close()
	// luego de creada el mock de la DB le decimos como va a responder cuando la invoquen
	mock.ExpectExec("BEGIN").
		WithArgs("USRDUMMY", "1.1.1.1", "out", "out", "out").
		WillReturnResult(sqlmock.NewResult(1, 1))
	// en la sentencia anterior le dijimos a nuestro mock que cada vez que alguien la invoque con la sentencia que inicia
	// con la palabra "begin" (tal cual como la llamamos desde el codigo productivo), y que si se hace con 5 argumentos
	// como "USRDUMMY", "1.1.1.1", entoces que responda de forma correcta un resultado sql

	//ahora teniendo nuestra DB mock preparada procedemos a crear el repo
	repo := NewRepository(db)
	// los cursores no se mockean por lo que en este caso solo podremos testear el error al ejecutar el llamado a la db
	// antes de testearlo inicializamos nuestros valores esperados
	stringEsperado := "" // string vacio que debe retornarse cuando se genera error al invocar a la db
	// ahora invoco a mi repo y obtengo sus respuestas
	stringObtenido, errObtenido := repo.LlamarSP()

	// ahora se hacen las validaciones
	// valido que el err devuelto sea diferente de nulo
	assert.NotNil(t, errObtenido)
	// valido que el string obtenido sea igual al esperado
	assert.Equal(t, stringEsperado, stringObtenido)

}
