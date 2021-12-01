package service

import (
	"github.com/nelbermora/go-interfaces/internal/repository"
)

// al igual que el repo se implementa el service a traves de interfaz
// recordemos los tres elementos son : Interface, struct y funcion constructora
// Interface notar que inicia con mayuscula para que sea publica
type Service interface {
	DummyFunc() (string, error)
}

// struct, aca definimos lo que necesitamos inyectarle al service, el repo por ejemplo
type service struct {
	repo repository.RepositoryInterfaz
}

// funcion constructora
func NewService(repoInyectado repository.RepositoryInterfaz) Service {
	return &service{
		repo: repoInyectado,
	}
}

// en este pkg puede ir la logica de negocio

func (s *service) DummyFunc() (string, error) {
	// para este ejemplo la funcion solo ejecutara la llamada al SP,
	// si esta se ejecuta correctamente entonces retornara un mensajem si no el error
	// ahora el repositorio se implementa a través de una interfaz, para invocar al repo debo crearlo con las dependencias que se requieran
	//return repository.LlamarSP()
	// esto debe hacerse al inicializarse la aplicacion, hacerlo aqui es una muy mala practica, pero a efectos del ejemplo lo haremos en este bloque de codigo
	// creo el repo indicando cual DB va a utilizar, le paso por parametro la base ya inicializada
	// se comenta esta bloque ya que ahora no se requiere construir aca el repo, pues ya viene inyectado cuando se crea el service
	/*
		myRepo := repository.NewRepository(db.MyDB)
		return myRepo.LlamarSP()
	*/
	return s.repo.LlamarSP()

}
