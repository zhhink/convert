package convert

import (
	"reflect"
	"unsafe"
)

// StringToBytes convert given string to []byte
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// StringToBytesSimple convert given string to []byte
func StringToBytesSimple(s string) []byte {
	return []byte(s)
}

// StringToMap convert s to map[string]interface{}
func StringToMap(s string) (m map[string]interface{}, err error) {
	jsonBytes := StringToBytes(s)
	return BytesToMap(jsonBytes)
}

// StringFillStruct fill a struct type
func StringFillStruct(str string, s interface{}) error {
	jsonBytes := StringToBytes(str)
	return BytesFillStruct(jsonBytes, &s)
}
