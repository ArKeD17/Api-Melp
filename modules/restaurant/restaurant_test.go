package restaurant

import (
	"reflect"
	"testing"
)

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
