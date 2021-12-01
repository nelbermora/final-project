package db

// dentro del folder clients pueden colocarse
// todos los clientes o librerias que usamos para invocar otras instancias
import (
	"database/sql"
	"fmt"
	"time"

	"github.com/godror/godror"
)

var MyDB *sql.DB

func InitializeDB() error {
	// esta funcion inicializa de forma automatica un pool de conexiones a la DB Oracle
	// se hace uso del driver godror que el la libreria oficial
	var P godror.ConnectionParams
	// los parametros de conexion deben ser buuscados en los secrets y/o variables de entorno
	// a efectos de este ejemplo se defienen en el codigo
	P.Username, P.Password = "DUMMYUSER", godror.NewPassword("DUMMYPASS")
	P.ConnectString = "host:port/ServiceName"
	P.SessionTimeout = 42 * time.Second
	fmt.Println(P.String())
	MyDB = sql.OpenDB(godror.NewConnector(P))
	err := MyDB.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
