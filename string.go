package ptr

import (
	"fmt"
	"reflect"
)

// S returns the string representation of the underlying value of the pointer.
// If the pointer is nil, S(nil) returns "<nil>".
func S(v interface{}) string {
	if v == nil {
		return fmt.Sprint(nil)
	}

	item := reflect.ValueOf(v)
	if item.Kind() == reflect.Ptr {
		item = item.Elem()
	}

	return fmt.Sprint(item)
}
