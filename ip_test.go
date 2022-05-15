package goip

import "testing"

func TestGetOutsideIp(t *testing.T) {
	t.Log(GetOutsideIp())
}

func TestGetInsideIp(t *testing.T) {
	t.Log(GetInsideIp())
}

func TestIps(t *testing.T) {
	t.Log(Ips())
}

func TestGetMacAddr(t *testing.T) {
	t.Log(GetMacAddr())
}
