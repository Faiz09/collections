package methods

import "reflect"

func makeSlice(i interface{}) reflect.Value {
	elemType := reflect.TypeOf(i).Elem()
	return reflect.MakeSlice(reflect.SliceOf(elemType), 0, 10)
}
