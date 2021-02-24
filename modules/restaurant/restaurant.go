package restaurant

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/mvochoa/logger"
	"gitlab.com/melp/api/libs/database"
	"gitlab.com/melp/api/modules"
)

// FIELDS campos de la tabla de restaurantes de la base de datos
const FIELDS = "id, rating, name, site, email, phone, street, city, state, lat, lng"

var (
	// ErrNotExistID no existe ningún registro con el ID
	ErrNotExistID = func(id interface{}) error {
		return fmt.Errorf("No existe ningún registro con el ID: '%v'", id)
	}
)

// Restaurants obtiene la lista de los restaurantes
func (r Restaurant) Restaurants(p graphql.ResolveParams) (interface{}, error) {
	result := ListRestaurants{}
	db := database.ConnectRoot()
	rows, err := db.Query(
		fmt.Sprintf(`
		SELECT %s
		FROM RESTAURANTS
		`, FIELDS),
	)

	if err == nil {
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(scan(&r)...)
			result.Data = append(result.Data, r)
			logger.Error("Restaurant:Restaurants - Scan", err)
		}
	} else if err.Error() != "sql: no rows in result set" {
		logger.Error("Restaurant:Restaurant", err)
		err = modules.ErrServer
	}

	logger.Error("Restaurant:Restaurants - CLOSE DATABASE", db.Close())
	return result, err
}

// GetRestaurant retorna los datos de un restaurante a partir del ID
func GetRestaurant(id string) (Restaurant, error) {
	var r Restaurant

	db := database.ConnectRoot()
	err := db.QueryRow(
		fmt.Sprintf(`
		SELECT %s
		FROM RESTAURANTS
		WHERE id = $1
		`, FIELDS),
		id,
	).Scan(scan(&r)...)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			err = ErrNotExistID(id)
		} else {
			logger.Error("Restaurant:GetRestaurant", err)
			err = modules.ErrServer
		}
	}

	logger.Error("Restaurant:GetRestaurant - CLOSE DATABASE", db.Close())
	return r, err
}

func scan(r *Restaurant) []interface{} {
	r.Email = &modules.EmailScalar{}
	r.Phone = &modules.PhoneScalar{}

	return []interface{}{&r.ID, &r.Rating, &r.Name, &r.Site, &r.Email.Value, &r.Phone.Value, &r.Street, &r.City, &r.State, &r.Lat, &r.Lng}
}
