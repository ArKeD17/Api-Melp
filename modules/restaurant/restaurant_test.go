package restaurant

import (
	"reflect"
	"testing"
)

func TestGetRestaurant(t *testing.T) {
	x := MockCreate()
	schema.DoNotToken(t, `{
		getRestaurant(id: "%s") {
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
	}`, x.ID)
	schema.ExistErrors(false)

	var result Restaurant
	schema.MapToStruct(MockParseMap(schema.GetData("getRestaurant")), &result, reflect.TypeOf(result))
	schema.AssertStruct(x, result)
}

func TestRestaurants(t *testing.T) {
	MockTruncate()
	x := ListRestaurants{
		Data: []Restaurant{MockCreate()},
	}
	r := schema.DoNotToken(t, `{
		restaurants {
			data {
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
		}
	}`)
	schema.ExistErrors(false)

	var result ListRestaurants
	restaurants := r.Data["restaurants"].(map[string]interface{})
	for k, v := range restaurants["data"].([]interface{}) {
		restaurants["data"].([]interface{})[k] = MockParseMap(v)
	}
	schema.MapToStruct(restaurants, &result, reflect.TypeOf(result))
	schema.AssertStruct(x, result)
}

func TestStatisticsRestaurants(t *testing.T) {
	x := MockCreate()
	schema.DoNotToken(t, `{
			statisticsRestaurants(lat: "%s", lng: "%s", radius: %d) {
				 count,
				 avg
				 std
			}
			}`, x.Lat, x.Lng, 80)
	schema.ExistErrors(false)

	var result StatisticsRestaurant
	schema.MapToStruct(schema.GetData("statisticsRestaurants"), &result, reflect.TypeOf(result))
}
