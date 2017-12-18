package gtmesh

import (
	"reflect"
	"testing"
)

func TestBufferRequest(t *testing.T) {
	type args struct {
		serviceName string
		payLoad     []byte
	}

	//Setup before test
	testNode.SetLocalHost("tcp://localhost:7072", "")
	// Setup finished here

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"no service and data", args{serviceName: "", payLoad: []byte{}}, []byte{}, true},
		{"echo but no data", args{serviceName: "echo", payLoad: []byte{}}, []byte{}, false},
		{"unregistered service", args{serviceName: "unknown", payLoad: []byte{}}, []byte{}, true},
		{"echo with data", args{serviceName: "echo", payLoad: []byte("testEcho")}, []byte("testEcho"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BufferRequest(tt.args.serviceName, tt.args.payLoad)
			if (err != nil) != tt.wantErr {
				t.Errorf("BufferRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err == nil && len(tt.args.payLoad) != 0) && (!reflect.DeepEqual(got, tt.want)) {
				t.Errorf("BufferRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
