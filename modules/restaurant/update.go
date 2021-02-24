package restaurant

import (
	"reflect"

	"github.com/graphql-go/graphql"
	"github.com/mvochoa/logger"
	"gitlab.com/melp/api/libs/database"
	"gitlab.com/melp/api/modules"
)

// Update actualizar los datos de un restaurante en especifico
func (r Restaurant) Update(p graphql.ResolveParams) (interface{}, error) {
	var err error
	r, err = GetRestaurant(p.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	err = modules.FillStruct(p.Args, &r, reflect.TypeOf(r))
	if err != nil {
		logger.Error("Restaurant:Update - FILL STRUCT", err)
		return nil, modules.ErrServer
	}

	db := database.ConnectRoot()
	result, err := db.Exec(`UPDATE RESTAURANTS SET rating = $2, name = $3, site = $4, email = $5, phone = $6, street = $7, city = $8, state = $9, lat = $10, lng = $11 WHERE id = $1`, r.ID, r.Rating, r.Name, r.Site, r.Email.Value, r.Phone.Value, r.Street, r.City, r.State, r.Lat, r.Lng)
	if err == nil {
		_, err = result.RowsAffected()
	}

	logger.Error("Restaurant:Update - CLOSE DATABASE", db.Close())
	if err == nil {
		return GetRestaurant(r.ID)
	}

	logger.Error("Restaurant:Update", err)
	return nil, modules.ErrServer
}
