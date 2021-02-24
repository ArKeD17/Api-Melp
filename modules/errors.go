package modules

import (
	"errors"
	"fmt"
)

// ErrServer error generico de un error no esperado
var ErrServer = errors.New("Hubo un error interno en el servidor")

// ErrNotExistFile retorna el error de cuando no existe el archivo
var ErrNotExistFile = func(name string) error {
	return fmt.Errorf("No existe el archivo '%s'", name)
}
