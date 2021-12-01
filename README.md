# go-testing

Una de las ventajas de implementar interfaces, es que se pueden "mockear" capas, esto permite desarrollar tests unitarios. Por ejemplo: el repo se contruye a través de una funcion constructora que recibe una Base de datos previamente inicializada. Entonces para testear la capa repositorio podemos crear un repository a través de la función constructora, pero inyectándole una capa "mockeada", es decir preparada por nosotros para que responde lo que queremos, a fin de cuentas lo que queremos probar es el repository, no la base de datos.

Para testear haremos uso de:

### 1. Package testing
Es una libreria estandar de Golang que permite crear testcases y validar respuestas

### 2. Package testify
Es una libreria de terceros ya homologada, que permite facilitar las validaciones esperadas.

## Ejecutando tests
Los tests serán ejecutados de forma automatica durante los pipelines de deploy. Pero claro está debemos garantizar que los tests estan ok ejecutandolos en nuetro entorno local a través del siguiente comando :
`$ go test`

Considerar que para el compilador reconozca un test, debe estar creado con el mismo nombre del archivo a probar concatenado con "_ test" y preferiblemente estar en el mismo folder. por ejemplo: 

Para probar el file repository.go creamos el test en un archivo llamado repository_test.go.

Un test podria verse de esta forma:

```go
package repository

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)
// es una buena practica iniciar cada test con el nombre Test
// cada test es una funcion que recibe por parametro una referencia del tipo testing.T
func TestDummy(t *testing.T) {
	valorEsperado := 3
    // suponiendo que vamos a pribar una funcion suma
    valorObtenido := sumar(1 + 2)
    // el paquete assert nos ahorra mucho codigo para hacer la validacion
    // con assert se pueden hacer distintas validaciones
    // Equal, NotEqual, Nil.. entre muchas otras
    // se puede documentar la documentacion oficial de testify si se desea conocer mas
    // en este caso se invoca el metodo Equal, que recibe, el test t, y los valores obtenido y esperado
    // si son igual es test estara ok, si no fallara el test.
	assert.Equal(t, valorObtenido, valorEsperado)    
}
```

