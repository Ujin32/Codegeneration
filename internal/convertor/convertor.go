package convertor

import (
	"reflect"
)

func MapToStruct(mp map[string]interface{}, item interface{}) error {
	valueStruct := reflect.ValueOf(item)
	elem := valueStruct.Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i)
		keyname := field.Tag.Get("keyname")
		if keyname == "" {
			keyname = field.Name
		}
		if val, ok := mp[keyname]; ok {
			fieldVal := reflect.ValueOf(val)
			elem.Field(i).Set(fieldVal)
		}

	}
	return nil
}

func StructToMap(item interface{}) map[string]interface{} {
	valuesStruct := reflect.ValueOf(item)
	if valuesStruct.Kind() != reflect.Struct {
		return nil
	}
	NumField := valuesStruct.Type().NumField()
	convertMap := make(map[string]interface{}, valuesStruct.Type().NumField())

	for i := 0; i < NumField; i++ {
		field := valuesStruct.Type().Field(i)
		keyname := field.Tag.Get("keyname")
		if keyname == "" {
			keyname = field.Name
		}

		if field.Type.Kind() == reflect.Struct {
			convertMap[keyname] = valuesStruct.Field(i).Interface()

		} else {
			convertMap[keyname] = valuesStruct.Field(i).Interface()
		}

	}
	return convertMap

}
