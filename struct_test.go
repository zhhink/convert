package convert

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type args struct {
		s interface{}
	}

	var funcArgs args
	funcArgs.s = testStructData

	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "struct to map test",
			args: funcArgs,
			want: testMap,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructToJSONBytes(t *testing.T) {
	type args struct {
		s interface{}
	}

	var funcArgs args
	funcArgs.s = testStructData

	wantBytes, _ := json.Marshal(testStructData)

	tests := []struct {
		name    string
		args    args
		wantB   []byte
		wantErr bool
	}{
		{
			name:    "TestStructToJSONBytes",
			args:    funcArgs,
			wantB:   wantBytes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := StructToJSONBytes(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToJSONBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("StructToJSONBytes() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestStructToJSONString(t *testing.T) {
	type args struct {
		s interface{}
	}

	var funcArgs args
	funcArgs.s = testStructData

	tests := []struct {
		name    string
		args    args
		wantStr string
		wantErr bool
	}{
		{
			name:    "TestStructToJSONString",
			args:    funcArgs,
			wantStr: testString,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := StructToJSONString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToJSONString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStr != tt.wantStr {
				t.Errorf("StructToJSONString() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}
