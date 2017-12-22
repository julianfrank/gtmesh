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
		{"empty service.ServiceName", testNode, args{service: "", handler: echoHandler}, true},
		{"empty handler", testNode, args{service: "x", handler: nil}, true},
		{"x,bufferHandler", testNode, args{service: "x", handler: echoHandler}, false},
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

func TestNode_addService(t *testing.T) {
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
		{"empty service.ServiceName", testNode, args{service: "", tcp: "x"}, true},
		{"empty tcp", testNode, args{service: "x", tcp: ""}, true},
		{"x,tcp", testNode, args{service: "x", tcp: "tcp"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.node.addService(tt.args.service, tt.args.tcp)

			if (err != nil) != tt.wantErr {
				t.Errorf("Node.addService() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if tt.node.ServiceStore[tt.args.service] == nil {
					t.Errorf("addService(service:%s,tcp:%s) Error:service not stored in ServiceStore %#v", tt.args.service, tt.args.tcp,tt.node.ServiceStore)
				}else if _,ok:=tt.node.ServiceStore[tt.args.service][tt.args.tcp]; !ok {
					t.Errorf("addService(service:%s,tcp:%s) Error:tcp not stored in ServiceStore %#v", tt.args.service, tt.args.tcp,tt.node.ServiceStore)
				}
				
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
		{"no service and data", testNode, args{serviceName: "", payLoad: []byte{}}, []byte{}, true},
		{"echo but no data", testNode, args{serviceName: "sys.echo", payLoad: []byte{}}, []byte{}, false},
		{"unregistered service", testNode, args{serviceName: "unknown", payLoad: []byte{}}, []byte{}, true},
		{"echo with data", testNode, args{serviceName: "sys.echo", payLoad: []byte("testEcho")}, []byte("testEcho"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testNode.SetLocalHost("tcp://localhost:7073", "")
			testNode.StartTCPServer()

			got, err := tt.node.BufferRequest(tt.args.serviceName, tt.args.payLoad)

			if (err != nil) != tt.wantErr {
				t.Errorf("Node.BufferRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err == nil && len(tt.args.payLoad) != 0) && (!reflect.DeepEqual(got, tt.want)) {
				t.Errorf("BufferRequest() = %v, want %v", got, tt.want)
			}

		})
	}
}
