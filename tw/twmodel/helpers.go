package twmodel

import (
	"reflect"
)

func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
