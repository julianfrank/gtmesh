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
		{"x,y", args{"x", "y"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetLocalHost(tt.args.tcp, tt.args.ws); (err != nil) != tt.wantErr {
				t.Errorf("SetLocalHost() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (LocalHost.TCPUrl != tt.args.tcp) || (LocalHost.WSUrl != tt.args.ws) {
				t.Errorf("SetLocalHost(tcp=%s,ws=%s)\tWanted:LocalHost{TCPurl:%s,WSurl:%s}\tGot:%t#v", tt.args.tcp, tt.args.ws, tt.args.tcp, tt.args.ws, LocalHost)
			}
		})
	}
}
