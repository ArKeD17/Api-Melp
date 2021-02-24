package restaurant

import (
	"testing"
)

func TestDeleteRestaurant(t *testing.T) {
	r := schema.DoNotToken(t, `mutation{
		deleteRestaurant(id: "%s")
	}`, MockCreate().ID)

	schema.ExistErrors(false)
	schema.Assert("Se elimino el restaurante con exito", r.Data["deleteRestaurant"].(string))
}
