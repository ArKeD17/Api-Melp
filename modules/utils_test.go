package modules

import (
	"reflect"
	"testing"
)

type StructForFillStruct struct {
	StringValue    string        `json:"string_value"`
	StringArray    []string      `json:"string_array"`
	IntValue       int           `json:"int_value"`
	IntArray       []int         `json:"int_array"`
	InterfaceArray []interface{} `json:"interface_array"`
}

func TestFillStruct(t *testing.T) {
	st := StructForFillStruct{
		StringValue:    "Value",
		IntValue:       1,
		StringArray:    []string{"value 1", "value 2"},
		IntArray:       []int{1, 2},
		InterfaceArray: []interface{}{"Value", 1, 1.2},
	}
	mp := map[string]interface{}{
		"string_value":    st.StringValue,
		"int_value":       st.IntValue,
		"string_array":    st.StringArray,
		"int_array":       st.IntArray,
		"interface_array": st.InterfaceArray,
	}

	var result StructForFillStruct
	err := FillStruct(mp, &result, reflect.TypeOf(result))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(st, result) {
		t.Errorf("Se esperaba que >>%+v<< fuera igual a >>%+v<<", result, st)
	}
}

func TestFillStructOfInterface(t *testing.T) {
	st := StructForFillStruct{
		StringValue:    "Value",
		IntValue:       1,
		StringArray:    []string{"value 1", "value 2"},
		IntArray:       []int{1, 2},
		InterfaceArray: []interface{}{"Value", 1, 1.2},
	}
	mp := map[string]interface{}{
		"string_value":    st.StringValue,
		"int_value":       st.IntValue,
		"string_array":    []interface{}{"value 1", "value 2"},
		"int_array":       []interface{}{1, 2},
		"interface_array": st.InterfaceArray,
	}

	var result StructForFillStruct
	err := FillStruct(mp, &result, reflect.TypeOf(result))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(st, result) {
		t.Errorf("Se esperaba que >>%+v<< fuera igual a >>%+v<<", result, st)
	}
}
