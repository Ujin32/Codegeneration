package convertor

import (
	"reflect"
)

func MapToStruct(mp map[string]interface{}, item interface{}) error {
	//Рефлексивное значение структуры
	valueStruct := reflect.ValueOf(item)
	//Значение структуры где есть доступ к изменению полей
	elem := valueStruct.Elem()

	for i := 0; i < elem.NumField(); i++ {
		//Получение рефлексивной структуры поля для i поля
		field := elem.Type().Field(i)
		//получение тега от рефлексивной структуры поля
		keyname := field.Tag.Get("keyname")
		if keyname == "" {
			keyname = field.Name
		}
		if val, ok := mp[keyname]; ok {
			//Преобразование значения мапы в рефлексивное значение
			fieldVal := reflect.ValueOf(val)
			//установка преобразованного значения в поле структуры
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

		convertMap[keyname] = valuesStruct.Field(i).Interface()

	}
	return convertMap

}
