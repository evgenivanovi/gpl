package mw

import (
	"net"
	"testing"
)

func TestIsIPInSubnet(t *testing.T) {

	datas := []struct {
		ip     string
		subnet string
		expect bool
	}{
		{
			ip:     "192.168.1.42",
			subnet: "192.168.1.0/24",
			expect: true,
		},
		{
			ip:     "10.0.0.1",
			subnet: "192.168.1.0/24",
			expect: false,
		},
	}

	t.Parallel()

	for _, data := range datas {
		t.Run("IsIpInSubnet", func(t *testing.T) {
			actual := isIPInSubnet(net.ParseIP(data.ip), data.subnet)
			if actual != data.expect {
				t.Errorf(
					"Get() = '%v', want '%v', for ip='%s' and subnet='%s'",
					actual, data.expect, data.ip, data.subnet,
				)
			}
		})
	}

}
