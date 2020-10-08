package convert

import (
	"reflect"
	"testing"
)

func TestMapFillStruct(t *testing.T) {

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
			if err := MapFillStruct(testMap, &fillStruct); (err != nil) != tt.wantErr {
				t.Errorf("MapFillStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(fillStruct, testStructData) {
				t.Errorf("MapFillStruct() fillStruct = %v, wantErr %v", fillStruct, testStructData)
			}
		})
	}
}

func TestMapToJSONBytes(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantB   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := MapToJSONBytes(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToJSONBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("MapToJSONBytes() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestMapToJSONString(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := MapToJSONString(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToJSONString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("MapToJSONString() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
