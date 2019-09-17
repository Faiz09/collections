package methods

import (
	"reflect"
)

func Filter(ci collectionInterface, f func(in interface{}) interface{}) interface{} {

	if reflect.TypeOf(ci).Kind() != reflect.Slice {
		panic("Not allowed for " + reflect.TypeOf(ci).Kind().String())
	}

	slice := makeSlice(ci)

	for j:=0; j<reflect.ValueOf(ci).Len() ;j++  {
		fy := f(reflect.ValueOf(ci).Index(j).Interface())
		if fy.(bool) {
			slice = reflect.Append(slice, reflect.ValueOf(ci).Index(j))
		}
	}

	return slice.Interface()

}