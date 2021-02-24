package restaurant

import (
	"reflect"
	"testing"

	"gitlab.com/melp/api/modules"
)

func TestCreateRestaurant(t *testing.T) {
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
		Lat:    "19.4400570537131",
		Lng:    "-99.1270470974249",
	}
	schema.DoNotToken(t, `mutation{
		createRestaurant(id: "%s", rating: %d,
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
				   lng,
			   }
	}`, x.ID, x.Rating, x.Name, x.Site, x.Email, x.Phone, x.Street, x.City, x.State, x.Lat, x.Lng)
	schema.ExistErrors(false)

	var result Restaurant
	schema.MapToStruct(MockParseMap(schema.GetData("createRestaurant")), &result, reflect.TypeOf(result))
	schema.AssertStruct(x, result)
}
