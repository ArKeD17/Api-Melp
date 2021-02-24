package restaurant

import (
	"reflect"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/mvochoa/logger"
	"gitlab.com/melp/api/libs/database"
	"gitlab.com/melp/api/modules"
)

// Create registra un nuevo restaurante
func (r Restaurant) Create(p graphql.ResolveParams) (interface{}, error) {
	err := modules.FillStruct(p.Args, &r, reflect.TypeOf(r))
	if err != nil {
		logger.Error("Restaurant:Create - FILL STRUCT", err)
		return nil, modules.ErrServer
	}

	lat, _ := strconv.ParseFloat(r.Lat, 13)
	lng, _ := strconv.ParseFloat(r.Lng, 13)

	db := database.ConnectAdmin()
	err = db.QueryRow(
		`
		INSERT INTO RESTAURANTS(id, rating, name, site, email, phone, street, city, state, lat, lng)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
		`,
		r.ID,
		r.Rating,
		r.Name,
		r.Site,
		r.Email.Value,
		r.Phone.Value,
		r.Street,
		r.City,
		r.State,
		lat,
		lng,
	).Scan(&r.ID)

	logger.Error("Restaurant:Create - CLOSE DATABASE", db.Close())

	if err == nil {
		return GetRestaurant(r.ID)
	}

	logger.Error("Restaurant:Create", err)
	err = modules.ErrServer

	return nil, err
}
