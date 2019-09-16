package methods

import (
	"reflect"
)

func Where(ci collectionInterface, key string, value interface{}) interface{} {

	if reflect.TypeOf(ci).Kind() != reflect.Slice {
		panic("Not allowed for " + reflect.TypeOf(ci).Kind().String())
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return whereString(ci, key, value.(string))
	case reflect.Int:
		return whereInt(ci, key, value.(int))
	default:
		panic("Value of type " + reflect.TypeOf(value).Kind().String() + " is not supported yet")
	}
}

func whereString(ci collectionInterface, key string, value string) interface{} {
	slice := makeSlice(ci)

	for j:=0; j<reflect.ValueOf(ci).Len() ;j++  {

		if reflect.ValueOf(ci).Index(j).FieldByName(key).String() == value {
			slice = reflect.Append(slice, reflect.ValueOf(ci).Index(j))
		}
	}

	return slice.Interface()
}

func whereInt(ci collectionInterface, key string, value int) interface{} {
	slice := makeSlice(ci)

	for j:=0; j<reflect.ValueOf(ci).Len() ;j++  {

		if reflect.ValueOf(ci).Index(j).FieldByName(key).Int() == int64(value) {
			slice = reflect.Append(slice, reflect.ValueOf(ci).Index(j))
		}
	}

	return slice.Interface()
}