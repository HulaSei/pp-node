package inbound

import (
	"testing"

	"github.com/perfect-panel/ppanel-node/api/panel"
)

func TestShouldConfigureTLSForNativeProtocols(t *testing.T) {
	tests := []struct {
		name string
		info panel.NodeInfo
		want bool
	}{
		{
			name: "TUIC managed certificate without generic security field",
			info: panel.NodeInfo{Type: "tuic", Protocol: &panel.Protocol{CertMode: "dns"}},
			want: true,
		},
		{
			name: "Hysteria local certificate",
			info: panel.NodeInfo{Type: "hysteria", Protocol: &panel.Protocol{CertMode: "file"}},
			want: true,
		},
		{
			name: "no certificate mode",
			info: panel.NodeInfo{Type: "tuic", Protocol: &panel.Protocol{CertMode: "none"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldConfigureTLS(&tt.info); got != tt.want {
				t.Fatalf("shouldConfigureTLS() = %t, want %t", got, tt.want)
			}
		})
	}
}
