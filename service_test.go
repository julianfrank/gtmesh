package gtmesh

import (
	"testing"

	"github.com/rsms/gotalk"
)

func TestAddLocalService(t *testing.T) {
	type args struct {
		service Service
		handler gotalk.BufferReqHandler
	}
	var bufferHandler gotalk.BufferReqHandler
	bufferHandler = func(s *gotalk.Sock, op string, payload []byte) ([]byte, error) { return nil, nil }
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"empty service.ServiceName", args{service: Service{ServiceName: ""}, handler: bufferHandler}, true},
		{"empty handler", args{service: Service{ServiceName: "x"}, handler: nil}, true},
		{"x,bufferHandler", args{service: Service{ServiceName: "x"}, handler: bufferHandler}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddLocalService(tt.args.service, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("AddLocalService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
