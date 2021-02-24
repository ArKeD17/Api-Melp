package restaurant

import (
	"reflect"
	"testing"

	"gitlab.com/melp/api/modules"
)

func TestUpdateRestaurant(t *testing.T) {
	x := Restaurant{
		ID:     MockCreate().ID,
		Rating: 4,
		Name:   "Nombre Actualizado" + modules.MockRand(),
		Site:   "www.melp.com",
		Email:  modules.NewEmailScalar("correoActualizado" + modules.MockRand() + "@melp.com"),
		Phone:  modules.NewPhoneScalar(modules.MockRand()[:10]),
		Street: "CalleActualizada" + modules.MockRand(),
		City:   "CiudadActualizada" + modules.MockRand(),
		State:  "EstadoActualizada" + modules.MockRand(),
		Lat:    "-92.1270470974249",
		Lng:    "34.4400570537131",
	}
	schema.DoNotToken(t, `mutation{
		updateRestaurant(
					id: "%s", 
					rating: %d,
					name: "%s",
					site: "%s",
					email: "%s",
					phone: "%s",
					street: "%s",
					city: "%s",
					state: "%s",
					lat: "%s",
					lng: "%s") {
						id,
						rating,
						name,
						site,
						email,
						phone,
						street,
						city,
						state,
						lat,
						lng
			  		 }
	}`, x.ID, x.Rating, x.Name, x.Site, x.Email, x.Phone, x.Street, x.City, x.State, x.Lat, x.Lng)
	schema.ExistErrors(false)

	var result Restaurant
	schema.MapToStruct(MockParseMap(schema.GetData("updateRestaurant")), &result, reflect.TypeOf(result))
	schema.AssertStruct(x, result)
}
