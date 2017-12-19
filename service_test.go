package gtmesh

import (
	"reflect"
	"testing"

	"github.com/rsms/gotalk"
)

func TestNode_AddLocalService(t *testing.T) {
	type args struct {
		service string
		handler gotalk.BufferReqHandler
	}
	tests := []struct {
		name    string
		node    *Node
		args    args
		wantErr bool
	}{
		{"empty service.ServiceName", &testNode, args{service: "", handler: echoHandler}, true},
		{"empty handler", &testNode, args{service: "x", handler: nil}, true},
		{"x,bufferHandler", &testNode, args{service: "x", handler: echoHandler}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testNode.SetLocalHost("tcp://localhost:7070", "")

			err := tt.node.AddLocalService(tt.args.service, tt.args.handler)

			if (err != nil) != tt.wantErr {
				t.Errorf("Node.AddLocalService() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if testNode.LocalServiceStore[tt.args.service] == nil {
					t.Errorf("AddLocalService() Error: LocalServiceStore %#v not updated with handler%#v", testNode.LocalServiceStore, tt.args.handler)
				}
			}
		})
	}
}

func TestNode_AddService(t *testing.T) {
	type args struct {
		service string
		tcp     string
	}
	tests := []struct {
		name    string
		node    *Node
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.node.AddService(tt.args.service, tt.args.tcp); (err != nil) != tt.wantErr {
				t.Errorf("Node.AddService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNode_BufferRequest(t *testing.T) {
	type args struct {
		serviceName string
		payLoad     []byte
	}
	tests := []struct {
		name    string
		node    *Node
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.node.BufferRequest(tt.args.serviceName, tt.args.payLoad)
			if (err != nil) != tt.wantErr {
				t.Errorf("Node.BufferRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.BufferRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
