package gtmesh

import (
	"testing"

	"github.com/rsms/gotalk"
)

func TestAddLocalService(t *testing.T) {
	type args struct {
		service string
		handler gotalk.BufferReqHandler
	}
	var bufferHandler gotalk.BufferReqHandler
	bufferHandler = func(s *gotalk.Sock, op string, payload []byte) ([]byte, error) { return nil, nil }
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"empty service.ServiceName", args{service: "", handler: bufferHandler}, true},
		{"empty handler", args{service: "x", handler: nil}, true},
		{"x,bufferHandler", args{service: "x", handler: bufferHandler}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddLocalService(tt.args.service, tt.args.handler)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddLocalService() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if LocalServiceStore[tt.args.service] == nil {
					t.Errorf("AddLocalService() Error: LocalServiceStore %#v not updated with handler%#v", LocalServiceStore, tt.args.handler)
				}
			}
		})
	}
}

func TestAddService(t *testing.T) {
	type args struct {
		service string
		tcp     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"empty service.ServiceName", args{service: "", tcp: "x"}, true},
		{"empty tcp", args{service: "x", tcp: ""}, true},
		{"x,tcp", args{service: "x", tcp: "tcp"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddService(tt.args.service, tt.args.tcp)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddService() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				found := false
				for _, host := range ServiceStore[tt.args.service] {
					if host == tt.args.tcp {
						found = true
					}
				}
				if !found {
					t.Errorf("AddService(service:%#v,tcp:%s) Error:tcp not stored in ServiceStore", tt.args.service, tt.args.tcp)
				}
			}
		})
	}
}
