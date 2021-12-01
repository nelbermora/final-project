package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"
	"strconv"
)

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

	_, err := rs.db.ExecContext(ctx, `BEGIN FTD.pkg_ftd_preca_procrear.SP_CUR_PFS_PROCREAR(:1, :2, :3, :4, :5); END;`, "USRDUMMY", "1.1.1.1", sql.Out{Dest: &cursor},
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
