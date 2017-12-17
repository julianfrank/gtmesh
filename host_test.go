package gtmesh

import "testing"

/*

func (t *testing.T) {
	type args struct {
		host Host
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"none", args{Host{TCPUrl: "", WSUrl: ""}}, true},
		{"emptyTCP", args{Host{TCPUrl: "", WSUrl: "WSUrl"}}, true},
		{"emptyWS", args{Host{TCPUrl: "TCPUrl", WSUrl: ""}}, false},
		{"both", args{Host{TCPUrl: "TCPUrl", WSUrl: "WSUrl"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartServers(tt.args.host); (err != nil) != tt.wantErr {
				t.Errorf("StartServers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_startTCPServer(t *testing.T) {
	type args struct {
		tcpURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"empty tcpURL", args{tcpURL: ""}, true},
		{"valid tcpURL", args{tcpURL: "localhost:7070"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := startTCPServer(tt.args.tcpURL); (err != nil) != tt.wantErr {
				t.Errorf("startTCPServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_startWSServer(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := startWSServer(); (err != nil) != tt.wantErr {
				t.Errorf("startWSServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/

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
				if (LocalHost.TCPUrl != tt.args.tcp) || (LocalHost.WSUrl != tt.args.ws) {
					t.Errorf("SetLocalHost(tcp=%s,ws=%s)\tWanted:gtmesh.Host{TCPurl:%s,WSurl:%s}\tGot:%#v", tt.args.tcp, tt.args.ws, tt.args.tcp, tt.args.ws, LocalHost)
				}
			}

		})
	}
}

/*
func TestStartServers(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartServers(); (err != nil) != tt.wantErr {
				t.Errorf("StartServers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_startTCPServer(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := startTCPServer(); (err != nil) != tt.wantErr {
				t.Errorf("startTCPServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_startWSServer(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := startWSServer(); (err != nil) != tt.wantErr {
				t.Errorf("startWSServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/
