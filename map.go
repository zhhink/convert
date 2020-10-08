package convert

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// MapFillStruct fill a struct type by map[string]interface{}
func MapFillStruct(m interface{}, s interface{}) error {
	// if !structs.IsStruct(s) {
	// 	return errors.New("type error, a struct type should be given")
	// }
	if err := mapstructure.Decode(m, &s); err != nil {
		return err
	}
	return nil
}

// MapToJSONBytes convert given map[string]interface{} to json []byte
func MapToJSONBytes(m interface{}) (b []byte, err error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// MapToJSONString convert given map[string]interface{} to json string
func MapToJSONString(m interface{}) (s string, err error) {
	jsonBytes, err := MapToJSONBytes(m)
	if err != nil {
		return "", err
	}
	return ByteSliceToString(jsonBytes), nil
}
