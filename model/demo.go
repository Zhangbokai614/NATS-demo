package model

import (
	"fmt"
	"reflect"
	"strconv"
)

type Payload interface {
	Bind(map[string]interface{})
}

type Cat struct {
	Color string `json:"color"`
	Age   int    `json:"age"`
}

// func (o *Cat) Bind(data map[string]interface{}) {
// 	o.Color = data["color"].(string)

// 	age, _ := strconv.Atoi(fmt.Sprintf("%1.0f", data["age"].(float64)))
// 	o.Age = age
// }

func (o *Cat) Bind(data map[string]interface{}) {
	bindValue(o, data)
}

func bindValue(payload interface{}, data map[string]interface{}) {
	v := reflect.ValueOf(payload).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)

		// get field tag name
		fieldTagName := field.Tag.Get("json")

		if fieldTagName == "" {
			fieldTagName = field.Name
		}

		// get field type
		fieldType := field.Type.String()

		// set value by type
		switch fieldType {
		case "string":
			v.Field(i).SetString(data[fieldTagName].(string))
		case "int":
			value, _ := strconv.Atoi(fmt.Sprintf("%1.0f", data[fieldTagName].(float64)))
			v.Field(i).SetInt(int64(value))

			// other type
		}
	}
}
