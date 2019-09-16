package methods

import (
	"reflect"
)

func Exists(ci collectionInterface, key string, value interface{}) bool {

	if reflect.TypeOf(ci).Kind() != reflect.Slice {
		panic("Not allowed for " + reflect.TypeOf(ci).Kind().String())
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		 return existsString(ci, key, value.(string))
	case reflect.Int:
		 return existsInteger(ci, key, value.(int))
	default:
		panic("Value of type " + reflect.TypeOf(key).Kind().String() + " is not supported yet")
	}
}

func existsString(ci collectionInterface, key string, value string) bool {

	for j:=0; j<reflect.ValueOf(ci).Len() ;j++  {

		if reflect.ValueOf(ci).Index(j).FieldByName(key).String() == value  {
			return true
		}
	}

	return false
}

func existsInteger(ci collectionInterface, key string, value int) bool {

	for j:=0; j<reflect.ValueOf(ci).Len() ;j++  {

		if int(reflect.ValueOf(ci).Index(j).FieldByName(key).Int()) == value  {
			return true
		}
	}

	return false
}