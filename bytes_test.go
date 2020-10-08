package convert

import (
	"reflect"
	"testing"
)

func TestBytesFillStruct(t *testing.T) {
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
			if err := BytesFillStruct(testBytes, &fillStruct); (err != nil) != tt.wantErr {
				t.Errorf("BytesFillStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(fillStruct, testStructData) {
				t.Errorf("fill struct = %v, want %v", fillStruct, testStructData)
			}
		})
	}
}

func TestBytesToMap(t *testing.T) {
	type args struct {
		b []byte
	}

	var funcArgs args
	funcArgs.b = testBytes

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
			_, err := BytesToMap(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}

	var funcArgs args
	funcArgs.b = testBytes

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: funcArgs,
			want: testString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteSliceToString(t *testing.T) {
	type args struct {
		bs []byte
	}

	var funcArgs args
	funcArgs.bs = testBytes

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: funcArgs,
			want: testString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteSliceToString(tt.args.bs); got != tt.want {
				t.Errorf("ByteSliceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToStringSimple(t *testing.T) {
	type args struct {
		b []byte
	}

	var funcArgs args
	funcArgs.b = testBytes

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: funcArgs,
			want: testString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToStringSimple(tt.args.b); got != tt.want {
				t.Errorf("ByteToStringSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}
