package gtmesh

import (
	"testing"
)

func TestSetLocalHost(t *testing.T) {
	type args struct {
		tcp string
		ws  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"empty tcp,ws", args{"", ""}, true},
		{"bad TCP and empty WS", args{"badtcp", ""}, true},
		{"valid TCP and empty WS", args{"tcp://localhost:7070", ""}, false},
		{"valid TCP and bad WS", args{"tcp://localhost:7070", "badws"}, true},
		{"valid TCP and WS but same port", args{"tcp://localhost:7070", "ws://localhost:7070"}, true},
		{"valid TCP and WS & different ports", args{"tcp://localhost:7070", "ws://localhost:7071"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetLocalHost(tt.args.tcp, tt.args.ws)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetLocalHost() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if LocalHost.TCPUrl == LocalHost.WSUrl {
					t.Errorf("SetLocalHost(tcp=%s,ws=%s)\tWanted:gtmesh.Host{TCPurl:%s,WSurl:%s}\tGot:%#v", tt.args.tcp, tt.args.ws, tt.args.tcp, tt.args.ws, LocalHost)
				}
			}

		})
	}
}

func Test_startTCPServer(t *testing.T) {

	tests := []struct {
		name    string
		tcp     string
		wantErr bool
	}{
		{"empty tcp", "", true},
		{"bad TCP", "badtcp", true},
		{"valid TCP", "tcp://localhost:7070", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := SetLocalHost(tt.tcp, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("startTCPServer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if err := startTCPServer(); (err != nil) != tt.wantErr {
					t.Errorf("startTCPServer() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func Test_startWSServer(t *testing.T) {
	tests := []struct {
		name    string
		ws      string
		wantErr bool
	}{

		{"empty ws", "", false},
		{"bad ws", "badws", true},
		{"valid ws", "ws://localhost:70898989871", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetLocalHost("tcp://localhost:7070", tt.ws)
			if (err != nil) != tt.wantErr {
				t.Errorf("startWSServer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if err := startWSServer(); (err != nil) != tt.wantErr {
					t.Errorf("startWSServer() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
