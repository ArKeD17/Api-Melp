package modules

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"unicode"
)

// Valid estructura para validacion de textos
type Valid struct {
	Text string
}

// Email valida si un correo electrónico es correcto
func (v Valid) Email() error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(v.Text) {
		return fmt.Errorf("El correo electrónico '%s' no es valido", v.Text)
	}

	return nil
}

// Password valida que una contraseña una longitud
// minima de 8 caracteres ademas de contener al menos un
// numero y una letra mayúscula
func (v Valid) Password() error {
	number := false
	upper := false
	for _, c := range v.Text {
		switch {
		case unicode.IsNumber(c):
			number = true
			break
		case unicode.IsUpper(c):
			upper = true
			break
		}
	}

	if !number || !upper || len(v.Text) < 8 {
		return errors.New("La contraseña debe tener una longitud minima de 8 caracteres ademas de contener al menos un numero y una letra mayúscula")
	}

	return nil
}

// Empty valida que una cadena tenga un longitud minima
func (v Valid) Empty(field string, length int) error {
	if len(v.Text) < length {
		return fmt.Errorf("El campo '%s' debe tener una longitud minima de %d caracteres", field, length)
	}

	return nil
}

// ImageExt valida que una cadena se el nombre de un archivo tipo imagen
func (v Valid) ImageExt() error {
	ext := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
		".svg":  true,
	}
	extFile := filepath.Ext(v.Text)
	if val, ok := ext[extFile]; !ok || !val {
		return fmt.Errorf("El archivo '%s' no es una imagen valida", v.Text)
	}

	return nil
}

// Phone valida un numero de telefono
func (v Valid) Phone() error {
	re := regexp.MustCompile(`^(\+\d{2})?[ -]?(\(\d{3}\))?[ -]?([ -]?\d{2}){5}$`)

	if !re.MatchString(v.Text) {
		return fmt.Errorf("El numero de telefono '%s' no es valido. Los formatos validos son: +001122334455, 1122334455, 11-22-33-44-55, 11 22 33 44 55, +00-11-22-33-44-55, +00 11 22 33 44 55", v.Text)
	}

	return nil
}

// Options valida si es una opción valida
func (v Valid) Options(options []string) error {
	for _, value := range options {
		if value == v.Text {
			return nil
		}
	}

	return fmt.Errorf("'%s' no es una opción valida", v.Text)
}
