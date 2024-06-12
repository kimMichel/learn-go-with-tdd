package reflection

import "reflect"

func traverse(x interface{}, fn func(input string)) {
	value := getValue(x)

	traverseValue := func(value reflect.Value) {
		traverse(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			traverseValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			traverseValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			traverseValue(value.MapIndex(key))
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}