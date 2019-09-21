package methods

import (
	"reflect"
	"sort"
)

func Sort(ci collectionInterface, key string) interface{} {

	sort.Slice(ci, func(i, j int) bool {
		switch reflect.ValueOf(ci).Index(i).FieldByName(key).Kind() {
		case reflect.String:
			return reflect.ValueOf(ci).Index(i).FieldByName(key).String() < reflect.ValueOf(ci).Index(j).FieldByName(key).String()
		case reflect.Int64:
			return reflect.ValueOf(ci).Index(i).FieldByName(key).Int() < reflect.ValueOf(ci).Index(j).FieldByName(key).Int()
		default:
			panic("value of type " + reflect.ValueOf(ci).Index(i).FieldByName(key).Kind().String() + " not supported!")
		}
	})

	return ci
}