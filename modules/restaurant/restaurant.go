package restaurant

import (
	"fmt"
	"math"
	"strconv"

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

// GetRestaurant consulta la información de una tarjeta en especifico
func (r Restaurant) GetRestaurant(p graphql.ResolveParams) (interface{}, error) {
	return GetRestaurant(p.Args["id"].(string))
}

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

// StatisticsRestaurants obtiene la lista de los restaurantes en base a distancia y radio en metros.
func (s StatisticsRestaurant) StatisticsRestaurants(p graphql.ResolveParams) (interface{}, error) {

	const R = 6371e3 // earth's mean radius in metres
	const pi = math.Pi

	lat, _ := strconv.ParseFloat(p.Args["lat"].(string), 64)
	lng, _ := strconv.ParseFloat(p.Args["lng"].(string), 64)
	radius := float64(p.Args["radius"].(int))

	minLat := lat - radius/R*180/pi
	maxLat := lat + radius/R*180/pi
	minLng := lng - radius/R*180/pi/math.Cos(lat*pi/180)
	maxLng := lng + radius/R*180/pi/math.Cos(lat*pi/180)

	db := database.ConnectRoot()
	err := db.QueryRow(
		fmt.Sprintf(`
		SELECT %s
		FROM RESTAURANTS WHERE lat Between $1 And $2 And lng Between $3 And $4
		`, "count(id) as count,(CASE WHEN avg(rating) IS NULL THEN 0 ELSE avg(rating) END) as avg, (CASE WHEN stddev(rating) IS NULL THEN 0 ELSE stddev(rating) END) as std"),
		minLat,
		maxLat,
		minLng,
		maxLng,
	).Scan(scanStatistics(&s)...)

	if err != nil {
		logger.Error("Restaurant:StatisticsRestaurants", err)
		err = modules.ErrServer
		return nil, err
	}

	logger.Error("Restaurant:StatisticsRestaurants - CLOSE DATABASE", db.Close())
	return s, nil
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

func scanStatistics(s *StatisticsRestaurant) []interface{} {

	return []interface{}{&s.Count, &s.Avg, &s.Std}
}
