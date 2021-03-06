package attr

import (
	"fmt"
	"reflect"
)

// IsEmpty check the value is empty or not
// empty: "", 0, nil,
func IsEmpty(v interface{}) bool {
	r := reflect.ValueOf(v)

	switch r.Kind() {
	case reflect.Int:
		i := v.(int)
		return IntIsEmpty(int64(i))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return IntIsEmpty(v.(int64))
	case reflect.Float32, reflect.Float64:
		return FloatIsEmpty(v.(float64))
	case reflect.String:
		return StringIsEmpty(v.(string))
	case reflect.Slice:
		return SliceIsEmpty(v)
	case reflect.Invalid:
		return InvalidIsEmpty(v)
	}

	panic(fmt.Sprintf("we have no case for (%v: %v) ", r.Kind(), v))
}

// IntIsEmpty check the int is empty or not
func IntIsEmpty(v int64) bool {
	if v == 0 {
		return true
	}
	return false
}

// StringIsEmpty check the string is empty or not
func StringIsEmpty(v string) bool {
	if v == "" {
		return true
	}
	return false
}

// FloatIsEmpty check the float is empty or not
func FloatIsEmpty(v float64) bool {
	if v > 0 || 0 < v {
		return false
	}
	return true
}

// SliceIsEmpty check the slice is empty or not
func SliceIsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.Slice:
		if r.Len() == 0 {
			return true
		}

		return false
	}

	panic(fmt.Sprintf("%v is not slice", v))
}

// InvalidIsEmpty check the invalid value is empty or not
func InvalidIsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	return false
}
