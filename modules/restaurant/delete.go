package restaurant

import (
	"github.com/graphql-go/graphql"
	"github.com/mvochoa/logger"
	"gitlab.com/melp/api/libs/database"
	"gitlab.com/melp/api/modules"
)

// Delete elimina un restaurante en especifico
func (r Restaurant) Delete(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	if _, err := GetRestaurant(id); err != nil {
		return nil, err
	}

	db := database.ConnectRoot()
	_, err := db.Exec(
		`
		DELETE FROM RESTAURANTS WHERE id = $1
		`,
		id,
	)

	if err != nil {
		logger.Error("Restaurant:Delete", err)
		err = modules.ErrServer
	}

	logger.Error("Restaurant:Delete - CLOSE DATABASE", db.Close())
	return "Se elimino el restaurante con exito", err
}
