package convert

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

// BytesFillStruct fill a stuct type by given []byte
func BytesFillStruct(b []byte, s interface{}) error {
	return json.Unmarshal(b, &s)
}

// BytesToMap convert given []byte to map[string]interface{}
func BytesToMap(b []byte) (m map[string]interface{}, err error) {
	mb := make(map[string]interface{})
	e := json.Unmarshal(b, &mb)
	return mb, e
}

// BytesToString convert given []byte to string
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

// ByteSliceToString convert given []byte to string
func ByteSliceToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// ByteToStringSimple convert []byte to string
func ByteToStringSimple(b []byte) string {
	return string(b)
}
