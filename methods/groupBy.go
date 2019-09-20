package methods

import (
	"reflect"
)

func GroupBy(ci collectionInterface, key string) interface{} {

	if reflect.TypeOf(ci).Kind() != reflect.Slice {
		panic("Not allowed for " + reflect.TypeOf(ci).Kind().String())
	}

	if reflect.ValueOf(ci).Len() < 1 {
		return reflect.ValueOf(ci)
	}

	t := reflect.ValueOf(ci).Index(0).FieldByName(key).Kind()

	switch t {
	case reflect.String:
		return stringValues(ci, key)
	case reflect.Int64:
		return intValues(ci, key)
	default:
		panic("value of type " + t.String() + " are not supported")
	}

}

func stringValues(ci collectionInterface, key string) interface{} {
	gMap := make(map[string][]interface{})

	for i:=0; i<reflect.ValueOf(ci).Len() ;i++  {
		mKey := reflect.ValueOf(ci).Index(i).FieldByName(key).String()
		gMap[mKey] = append(gMap[mKey], reflect.ValueOf(ci).Index(i).Interface())
	}

	return gMap
}

func intValues(ci collectionInterface, key string) interface{} {
	gMap := make(map[int64][]interface{})

	for i:=0; i<reflect.ValueOf(ci).Len() ;i++  {
		mKey := reflect.ValueOf(ci).Index(i).FieldByName(key).Int()
		gMap[mKey] = append(gMap[mKey], reflect.ValueOf(ci).Index(i).Interface())
	}

	return gMap
}