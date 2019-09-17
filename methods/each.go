package methods

import (
	"reflect"
)

func Each(ci collectionInterface, f func(in interface{}) interface{}) interface{} {

	if reflect.TypeOf(ci).Kind() != reflect.Slice {
		panic("Not allowed for " + reflect.TypeOf(ci).Kind().String())
	}

	slice := makeSlice(ci)

	for j:=0; j<reflect.ValueOf(ci).Len() ;j++  {
		fy := f(reflect.ValueOf(ci).Index(j).Interface())
		slice = reflect.Append(slice, reflect.ValueOf(fy))
	}

	return slice.Interface()

}