package restaurant

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mvochoa/logger"
	"gitlab.com/melp/api/libs/database"
	"gitlab.com/melp/api/modules"
)

var schema = modules.Graphql{Schema: func() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: Queries,
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: Mutations,
		}),
	})
	logger.Error("Restaurant:Testing", err)
	return schema
}()}

// MockCreate crea un restaurante
func MockCreate() Restaurant {
	x := Restaurant{
		ID:     "ID " + modules.MockRand(),
		Rating: 4,
		Name:   "Nombre " + modules.MockRand(),
		Site:   "www.melp.com",
		Email:  modules.NewEmailScalar("correo" + modules.MockRand() + "@melp.com"),
		Phone:  modules.NewPhoneScalar(modules.MockRand()[:10]),
		Street: "Calle" + modules.MockRand(),
		City:   "Ciudad" + modules.MockRand(),
		State:  "Estado" + modules.MockRand(),
		Lat:    "19.1270470974249",
		Lng:    "-99.4400570537131",
	}

	p := graphql.ResolveParams{}
	bytes, _ := json.Marshal(x)
	json.Unmarshal(bytes, &p.Args)

	p.Args["rating"] = int(p.Args["rating"].(float64))

	data, err := x.Create(p)
	logger.Error("Restaurant:MockCreate", err)
	return data.(Restaurant)
}

// MockParseMap ajusta el mapa de restaurantes
func MockParseMap(m interface{}) map[string]interface{} {
	if _, ok := m.(map[string]interface{}); !ok {
		logger.Fatal("Restaurant:MockParseMap", errors.New("La data no es un mapa"))
	}
	data := m.(map[string]interface{})
	data["email"] = modules.NewEmailScalar(data["email"].(string))
	data["phone"] = modules.NewPhoneScalar(data["phone"].(string))

	return data
}

// MockTruncate elimina todos los registros de la tabla de RESTAURANTS
func MockTruncate() {
	db := database.ConnectRoot()
	_, err := db.Exec(`TRUNCATE TABLE RESTAURANTS CASCADE`)
	logger.Fatal("Restaurant:MockTruncate", err)
	logger.Fatal("Restaurant:MockTruncate - CLOSE DATABASE", db.Close())
	time.Sleep(100 * time.Millisecond)
}
