package gtmesh

import "testing"

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
		{",y", args{"", "y"}, true},
		{"x,", args{"x", ""}, false},
		{"x,y", args{"x", "y"}, false},
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
