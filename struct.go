package convert

import (
	"encoding/json"

	"github.com/fatih/structs"
)

// StructToMap convert struct to a map[string]interface{}
func StructToMap(s interface{}) map[string]interface{} {
	return structs.Map(s)
}

// StructToJSONBytes convert given struct to a json []byte
func StructToJSONBytes(s interface{}) (b []byte, err error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// StructToJSONString convert given struct to a json string
func StructToJSONString(s interface{}) (str string, err error) {
	jsonBytes, err := StructToJSONBytes(s)
	if err != nil {
		return "", err
	}
	return ByteSliceToString(jsonBytes), nil
}
