package methods

import (
	"reflect"
)

func Avg(ci collectionInterface, key string) interface{} {

	if reflect.TypeOf(ci).Kind() != reflect.Slice {
		panic("Not allowed for " + reflect.TypeOf(ci).Kind().String())
	}

	switch reflect.TypeOf(key).Kind() {
	case reflect.String:
		 return avgString(ci, key)
	default:
		panic("Value of type " + reflect.TypeOf(key).Kind().String() + " is not supported yet")
	}
}

func avgString(ci collectionInterface, key string) int {

	sum := 0

	for j:=0; j<reflect.ValueOf(ci).Len() ;j++  {

		if reflect.ValueOf(ci).Index(j).FieldByName(key).Kind() == reflect.Int64 {
			sum += int(reflect.ValueOf(ci).Index(j).FieldByName(key).Int())
		} else {
			panic("Operation not supported for " + reflect.ValueOf(ci).Index(j).FieldByName(key).Kind().String() + " values")
		}
	}

	return sum / reflect.ValueOf(ci).Len()

}