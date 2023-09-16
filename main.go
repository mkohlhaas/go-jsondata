package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myObject struct {
	ArrayValue []int `json:"array_value"`
}

type myJSON struct {
	IntValue        int       `json:"int_value"`
	BoolValue       bool      `json:"bool_value"`
	StringValue     string    `json:"string_value"`
	DateValue       time.Time `json:"date_value"`
	ObjectValue     *myObject `json:"object_value"`
	NullStringValue *string   `json:"nullString_value,omitempty"`
	NullIntValue    *int      `json:"nullInt_value"`
	EmptyString     string    `json:"emptyString,omitempty"`
}

// /////////////////// Main ///////////////////////////////////////
func main() {
	// I. Marshalling/Serialization/Encoding

	// 1. Using a generic string-any-map
	data := map[string]any{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"now":         time.Now(),
		"dateValue":   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
		"nullStringValue": nil,
		"nullIntValue":    nil,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("%-13s%s\n", "json data:", jsonData)

	// 2. Using a struct
	otherInt := 4321
	data1 := &myJSON{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		ObjectValue: &myObject{
			ArrayValue: []int{1, 2, 3, 4},
		},
		NullStringValue: nil,
		NullIntValue:    &otherInt,
	}

	jsonData1, err := json.Marshal(data1)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("%-13s%s\n", "json1 data:", jsonData1)

	// II. Unmarshalling/Deserialization/Decoding

	// 1. Using a generic string-any-map
	jsonData2 := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null,
      "extraValue":4321
		}
	`

	var data2 map[string]interface{}
	err = json.Unmarshal([]byte(jsonData2), &data2)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("%-13v%s\n", "json2 map:", data2)

	rawDateValue, ok := data2["dateValue"]
	if !ok {
		fmt.Printf("dateValue does not exist\n")
		return
	}

	// Using type assertion to access unmarshalled data
	dateValue, ok := rawDateValue.(string)
	if !ok {
		fmt.Printf("dateValue is not a string\n")
		return
	}
	fmt.Printf("%-13s%s\n", "date value:", dateValue)

	// 2. Using a struct

	var data3 *myJSON
	err = json.Unmarshal([]byte(jsonData), &data3)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("%-13s%#v\n", "json struct:", data3)
	fmt.Printf("%-13s%#v\n", "dateValue:", data3.DateValue)
	fmt.Printf("%-13s%#v\n", "objectValue:", data3.ObjectValue)
}
