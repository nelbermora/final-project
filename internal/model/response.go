package model

import "errors"

// en models pueden definirse los modelos, entidades o dominios
// son todos aquellos objetos o 'clases' que definen nuestra entidad principal de negocio y las complementarias
type Response struct {
	Message string `json:"message"`
}

type Request struct {
	Input string `json:"input"`
}

// se puede aprovechar y definir funciones para el objeto request que permita validar si el objeto esta correcto y valido
// esto es util si se quieren hacer validaciones en el request o casos similares
func (r *Request) Validate() error {
	if len(r.Input) < 4 {
		return errors.New("input incompleto")
	}
	return nil
}
