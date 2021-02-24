package modules

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/mvochoa/logger"
	"golang.org/x/crypto/bcrypt"
)

// GetNameFile copia el archivo temporal a la capeta destino
func GetNameFile(tmpName, prefixName string) (string, error) {
	path := os.Getenv("HOME") + "/tmp/" + tmpName
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		name := fmt.Sprintf("%s%d%s", prefixName, time.Now().UnixNano(), filepath.Ext(path))
		err = MoveFile(path, fmt.Sprintf("%s/storage/%s", os.Getenv("HOME"), name))
		if err != nil {
			logger.Error("GetNameFile - RENAME FILE", err)
			return "", ErrServer
		}

		return name, nil
	}

	return prefixName + "0.png", ErrNotExistFile(tmpName)
}

// MoveFile mueve un archivo de ubicación a otra
func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("No se pudo abrir el archivo fuente: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("No se pudo abrir el archivo destino: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Error al escribir en el archivo de salida: %s", err)
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Error al eliminar el archivo original: %s", err)
	}
	return nil
}

// HashPassword genera el hash a partir de la contraseña
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	logger.Error("HashPassword - CREATE HASH PASSWORD", err)
	return hex.EncodeToString(bytes)
}

// SetField asigna el valor a la estructura
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	fieldVal := structValue.FieldByName(name)

	if !fieldVal.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !fieldVal.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	if value == nil {
		return nil
	}

	val := reflect.ValueOf(value)

	if fieldVal.Type() != val.Type() {

		if m, ok := value.(map[string]interface{}); ok {
			if fieldVal.Kind() == reflect.Struct {
				return FillStruct(m, fieldVal.Addr().Interface(), fieldVal.Type())
			}

			if fieldVal.Kind() == reflect.Ptr && fieldVal.Type().Elem().Kind() == reflect.Struct {
				if fieldVal.IsNil() {
					fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
				}
				if fieldVal.Kind() == reflect.Ptr {
					return FillStruct(m, fieldVal.Interface(), fieldVal.Type().Elem())
				}
				return FillStruct(m, fieldVal.Interface(), fieldVal.Type())
			}

		} else if fieldVal.Kind() == reflect.Slice {
			arr := reflect.MakeSlice(reflect.SliceOf(fieldVal.Type().Elem()), 0, val.Cap())
			for _, el := range value.([]interface{}) {
				s := reflect.New(fieldVal.Type().Elem()).Interface()
				t := fieldVal.Type().Elem()
				if m, ok := el.(map[string]interface{}); ok {

					var field reflect.StructField
					tagJSON := make(map[string]string)
					for i := 0; i < t.NumField(); i++ {
						field = t.Field(i)
						tagJSON[field.Tag.Get("json")] = field.Name
					}

					var ok bool
					var err error
					for k, v := range m {
						if _, ok = tagJSON[k]; ok {
							err = SetField(s, tagJSON[k], v)
							if err != nil {
								return err
							}
						}
					}

				} else {
					arr = reflect.Append(arr, reflect.ValueOf(el))
					continue
				}

				arr = reflect.Append(arr, reflect.ValueOf(s).Elem())
			}

			fieldVal.Set(arr)
			return nil
		}

		return fmt.Errorf("Provided value type didn't match obj field type %s:%v Value:%v", name, fieldVal.Type(), val.Type())
	}

	fieldVal.Set(val)
	return nil

}

// FillStruct extrae los valores del mapa y los nombre de los tags de la estructura
func FillStruct(m map[string]interface{}, s interface{}, t reflect.Type) error {
	var field reflect.StructField
	tagJSON := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		field = t.Field(i)
		tagJSON[field.Tag.Get("json")] = field.Name
	}

	var ok bool
	var err error
	for k, v := range m {
		if _, ok = tagJSON[k]; ok {
			err = SetField(s, tagJSON[k], v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
