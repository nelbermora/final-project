package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"
	"strconv"
)

// comento el codigo para implementarlo con interfaz e inyeccion de la DB
/*
func LlamarSP() (string, error) {
	ctx := context.Background()
	var desc string
	var cod int
	var cursor driver.Rows
	// para este ejemplo existe un SP predeterminado que recibe 5 parametros de entrada
	// de los cuales lso ultimos tres los retorna en la salida tambien
	// el penultimo valor es un codigo de error.
	_, err := db.MyDB.ExecContext(ctx, `BEGIN Schema.Package.StoredProcedure(:1, :2, :3, :4, :5); END;`, "USRDUMMY", "1.1.1.1", sql.Out{Dest: &cursor},
		sql.Out{Dest: &cod}, sql.Out{Dest: &desc})
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer cursor.Close()
	if cod == 0 {
		return "SP invocado satisfactoriamente", nil
	}
	return "SP invocado. Codigo de retorno: " + strconv.Itoa(cod), nil

}*/

type RepositoryInterfaz interface {
	LlamarSP() (string, error)
}

type repositoryStruct struct {
	db *sql.DB
}

func NewRepository(dbAUtilizar *sql.DB) RepositoryInterfaz {
	return &repositoryStruct{
		db: dbAUtilizar,
	}
}

func (rs *repositoryStruct) LlamarSP() (string, error) {
	ctx := context.Background()
	var desc string
	var cod int
	var cursor driver.Rows

	//aca en vez de llamar directamente a la db, usamos la base de datos inyectada o definida cuando se construyo el repo
	//_, err := db.MyDB.ExecContext(ctx, `BEGIN Schema.Package.StoredProcedure(:1, :2, :3, :4, :5); END;`, "USRDUMMY", "1.1.1.1", sql.Out{Dest: &cursor},
	//	sql.Out{Dest: &cod}, sql.Out{Dest: &desc})
	_, err := rs.db.ExecContext(ctx, `BEGIN Schema.Package.StoredProcedure(:1, :2, :3, :4, :5); END;`, "USRDUMMY", "1.1.1.1", sql.Out{Dest: &cursor},
		sql.Out{Dest: &cod}, sql.Out{Dest: &desc})
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer cursor.Close()
	if cod == 0 {
		return "SP invocado satisfactoriamente", nil
	}
	return "SP invocado. Codigo de retorno: " + strconv.Itoa(cod), nil
}
