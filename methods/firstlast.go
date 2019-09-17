package methods

import (
	"reflect"
)

func First(ci collectionInterface) interface{} {

	if reflect.TypeOf(ci).Kind() != reflect.Slice {
		panic("Not allowed for " + reflect.TypeOf(ci).Kind().String())
	}

	return reflect.ValueOf(ci).Index(0).Interface()
}

func Last(ci collectionInterface) interface{} {

	if reflect.TypeOf(ci).Kind() != reflect.Slice {
		panic("Not allowed for " + reflect.TypeOf(ci).Kind().String())
	}

	return reflect.ValueOf(ci).Index(int(reflect.ValueOf(ci).Len()) - 1).Interface()
}