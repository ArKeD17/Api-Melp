package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mvochoa/logger"
)

// Graphql estructura para las consultas a graphql
type Graphql struct {
	Schema graphql.Schema
	Result GraphqlResult
	T      *testing.T
}

// GraphqlResult estructura para los resultado de graphql
type GraphqlResult struct {
	Errors []error
	Data   map[string]interface{}
}

// MockRand retorna un numero aleatorio en forma de cadena
func MockRand() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// Do ejecuta queries de graphql que necesitan un token
// Level User:
//	-1: EXTERNAL
//	0: ASISTENTE
// 	1: ADMINISTRADOR
// func (g *Graphql) Do(t *testing.T, idUser int, query string, args ...interface{}) GraphqlResult {

// 	token, _ := CreateToken(strconv.Itoa(idUser), 0)
// 	return g.exec(t, graphql.Do(graphql.Params{
// 		Schema:        g.Schema,
// 		RequestString: fmt.Sprintf(query, args...),
// 		Context:       context.WithValue(context.Background(), CtxToken{}, token),
// 	}))
// }

// DoNotToken ejecuta queries de graphql sin un token
func (g *Graphql) DoNotToken(t *testing.T, query string, args ...interface{}) GraphqlResult {
	return g.exec(t, graphql.Do(graphql.Params{
		Schema:        g.Schema,
		RequestString: fmt.Sprintf(query, args...),
	}))
}

// DoWithToken ejecuta queries de graphql con un token especifico
// func (g *Graphql) DoWithToken(t *testing.T, token, query string, args ...interface{}) GraphqlResult {
// 	return g.exec(t, graphql.Do(graphql.Params{
// 		Schema:        g.Schema,
// 		RequestString: fmt.Sprintf(query, args...),
// 		Context:       context.WithValue(context.Background(), CtxToken{}, token),
// 	}))
// }

func (g *Graphql) exec(t *testing.T, result *graphql.Result) GraphqlResult {
	g.Result = GraphqlResult{
		Data:   make(map[string]interface{}),
		Errors: make([]error, len(result.Errors)),
	}
	if result.Data != nil {
		g.Result.Data = result.Data.(map[string]interface{})
	}
	for k, v := range result.Errors {
		g.Result.Errors[k] = errors.New(v.Message)
	}

	g.T = t
	return g.Result
}

// ExistsFile valida si existe el archivo
func (g *Graphql) ExistsFile(path string) {
	g.existsFile(path, false)
}

// AssertStruct compara si dos estructuras son iguales
func (g *Graphql) AssertStruct(s1, s2 interface{}) {
	g.compareStruct(s1, s2, false)
}

// AssertErrors compara si dos errores son iguales
func (g *Graphql) AssertErrors(e1, e2 error) {
	g.compareErrors(e1, e2, false)
}

// Assert compara si dos errores son iguales
func (g *Graphql) Assert(v1, v2 interface{}) {
	g.compare(v1, v2, false)
}

// NotExistsFile valida que no existe el archivo
func (g *Graphql) NotExistsFile(path string) {
	g.existsFile(path, true)
}

// NotAssertStruct compara si dos estructuras NO son iguales
func (g *Graphql) NotAssertStruct(s1, s2 interface{}) {
	g.compareStruct(s1, s2, true)
}

// NotAssertErrors compara si dos errores NO son iguales
func (g *Graphql) NotAssertErrors(e1, e2 error) {
	g.compareErrors(e1, e2, true)
}

// NotAssert compara si dos errores NO son iguales
func (g *Graphql) NotAssert(v1, v2 interface{}) {
	g.compare(v1, v2, true)
}

func (g *Graphql) compareArray(a1, a2 []interface{}, not bool, prefix string) {
	msg := ""
	if not {
		msg = "NO "
	}

	if len(a1) != len(a2) {
		g.T.Errorf("%s %sSe esperaba que el arreglo \n\n'%+v::%T'\n\nsea del mismo tamaño de '%+v::%T'",
			prefix, msg, a1, a1, a2, a2)
		return
	}

	for i1 := range a1 {
		switch a1[i1].(type) {
		case map[string]interface{}:
			if _, ok := a2[i1].(map[string]interface{}); !ok {
				g.T.Errorf("%s[%v] %sSe esperaba '%+v::%T' pero se recibió '%+v::%T'",
					prefix, i1, msg, a1[i1], a1[i1], a2[i1], a2[i1])
				break
			}
			g.compareMap(a1[i1].(map[string]interface{}), a2[i1].(map[string]interface{}), not, fmt.Sprintf("%s[%v]", prefix, i1))
			break
		case []interface{}:
			if _, ok := a2[i1].([]interface{}); !ok {
				g.T.Errorf("%s[%v] %sSe esperaba '%+v::%T' pero se recibió '%+v::%T'",
					prefix, i1, msg, a1[i1], a1[i1], a2[i1], a2[i1])
				break
			}
			g.compareArray(a1[i1].([]interface{}), a2[i1].([]interface{}), not, fmt.Sprintf("%s[%v]", prefix, i1))
			break
		default:
			if reflect.DeepEqual(a1[i1], a2[i1]) == not {
				g.T.Errorf("%s[%v] %sSe esperaba '%+v::%T' pero se recibió '%+v::%T'",
					prefix, i1, msg, a1[i1], a1[i1], a2[i1], a2[i1])
			}
		}
	}
}

func (g *Graphql) compareMap(m1, m2 map[string]interface{}, not bool, prefix string) {
	msg := ""
	if not {
		msg = "NO "
	}
	for k1 := range m1 {
		switch m1[k1].(type) {
		case map[string]interface{}:
			if _, ok := m2[k1].(map[string]interface{}); !ok {
				g.T.Errorf("%s[%v] %sSe esperaba '%+v::%T' pero se recibió '%+v::%T'",
					prefix, k1, msg, m1[k1], m1[k1], m2[k1], m2[k1])
				break
			}
			g.compareMap(m1[k1].(map[string]interface{}), m2[k1].(map[string]interface{}), not, fmt.Sprintf("%s[%v]", prefix, k1))
			break
		case []interface{}:
			if _, ok := m2[k1].([]interface{}); !ok {
				g.T.Errorf("%s[%v] %sSe esperaba '%+v::%T' pero se recibió '%+v::%T'",
					prefix, k1, msg, m1[k1], m1[k1], m2[k1], m2[k1])
				break
			}
			g.compareArray(m1[k1].([]interface{}), m2[k1].([]interface{}), not, fmt.Sprintf("%s[%v]", prefix, k1))
			break
		default:
			if reflect.DeepEqual(m1[k1], m2[k1]) == not {
				g.T.Errorf("%s[%v] %sSe esperaba '%+v::%T' pero se recibió '%+v::%T'",
					prefix, k1, msg, m1[k1], m1[k1], m2[k1], m2[k1])
			}
		}
	}
}

func (g *Graphql) existsFile(path string, not bool) {
	m := ""
	if not {
		m = "NO "
	}

	_, err := os.Stat(path)
	if !os.IsNotExist(err) == not {
		g.T.Errorf("%sSe encontró el archivo '%s'", m, path)
	}
}

func (g *Graphql) compareStruct(s1, s2 interface{}, not bool) {
	b1, err := json.Marshal(s1)
	logger.Fatal("Testing:CompareStruct - MARSHAL 1", err)
	b2, err := json.Marshal(s2)
	logger.Fatal("Testing:CompareStruct - MARSHAL 2", err)

	var m1, m2 map[string]interface{}
	err = json.Unmarshal(b1, &m1)
	logger.Fatal("Testing:CompareStruct - UNMARSHAL 1", err)
	err = json.Unmarshal(b2, &m2)
	logger.Fatal("Testing:CompareStruct - UNMARSHAL 2", err)

	g.compareMap(m1, m2, not, "")
}

func (g *Graphql) compareErrors(e1, e2 error, not bool) {
	m := ""
	if not {
		m = "NO "
	}
	if reflect.DeepEqual(e1, e2) == not {
		g.T.Errorf("%sSe esperaba el error '%v' pero se recibió '%v'", m, e2, e1)
	}
}

func (g *Graphql) compare(v1, v2 interface{}, not bool) {
	m := ""
	if not {
		m = "NO "
	}
	if reflect.DeepEqual(v1, v2) == not {
		g.T.Errorf("%sSe esperaba '%v::%T' pero se recibió '%v::%T'", m, v2, v2, v1, v1)
	}
}

// ExistErrors valida si existen errores
func (g *Graphql) ExistErrors(wantErrors bool) {
	if len(g.Result.Errors) > 0 && !wantErrors {
		g.T.Errorf("Error al ejecutar la operación graphql, errores: %+v", g.Result.Errors)
	}

	if len(g.Result.Errors) == 0 && wantErrors {
		g.T.Errorf("No se recibió ningún error al ejecutar la operación graphql")
	}
}

// MapToStruct convierte un mapa en un estructura
func (g *Graphql) MapToStruct(data map[string]interface{}, pointerStruct interface{}, tp reflect.Type) {
	err := FillStruct(data, pointerStruct, tp)
	if err != nil {
		g.T.Errorf("Se esta devolviendo un error al transformar la data: %+v", err)
	}
}

// GetData extrae la data de un GraphqlResult
func (g *Graphql) GetData(key string) map[string]interface{} {
	var ok bool
	var data interface{}
	if data, ok = g.Result.Data[key]; !ok {
		g.T.Errorf("No se esta devolviendo ninguna data: %+v", g.Result.Data)
	}

	return data.(map[string]interface{})
}
