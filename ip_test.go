package goip

import "testing"

func TestIpIp(t *testing.T) {
	t.Log("GetOutsideIp", GetOutsideIp())
	t.Log("GetInsideIp", GetInsideIp())
	t.Log(Ips())
	t.Log("GetMacAddr", GetMacAddr())
}
