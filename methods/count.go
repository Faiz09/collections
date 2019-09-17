package methods

import (
	"reflect"
)

func Count(ci collectionInterface) int {
	return int(reflect.ValueOf(ci).Len())
}