package convert

import (
	"reflect"
	"testing"
)

type testStruct struct {
	Name  string
	Age   int
	Email string
}

var (
	testString     string
	testBytes      []byte
	testStructData testStruct
	testMap        map[string]interface{}
)

func init() {
	testString = `{"Name":"default user","Age":20,"Email":"default@163.com"}`
	testBytes = []byte(testString)
	testStructData = testStruct{
		Name:  "default user",
		Age:   20,
		Email: "default@163.com",
	}
	testMap = map[string]interface{}{
		"Name":  "default user",
		"Age":   20,
		"Email": "default@163.com",
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}

	var funcArgs args
	funcArgs.s = testString

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "TestStringToBytes",
			args: funcArgs,
			want: testBytes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytesSimple(t *testing.T) {
	type args struct {
		s string
	}

	var funcArgs args
	funcArgs.s = testString

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "",
			args: funcArgs,
			want: testBytes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytesSimple(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytesSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToMap(t *testing.T) {
	type args struct {
		s string
	}

	var funcArgs args
	funcArgs.s = testString

	tests := []struct {
		name    string
		args    args
		wantM   map[string]interface{}
		wantErr bool
	}{
		{
			name:    "",
			args:    funcArgs,
			wantM:   testMap,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := StringToMap(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStringFillStruct(t *testing.T) {

	var fillStruct testStruct

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StringFillStruct(testString, &fillStruct); (err != nil) != tt.wantErr {
				t.Errorf("StringFillStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(fillStruct, testStructData) {
				t.Errorf("fill struct = %v, want %v", fillStruct, testStructData)
			}
		})
	}
}
