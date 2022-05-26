package goip

import "testing"

func TestGetOutsideIp(t *testing.T) {
	t.Log(GetOutsideIp())
}

func BenchmarkGetOutsideIp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetOutsideIp())
	}
}

func TestGetInsideIp(t *testing.T) {
	t.Log(GetInsideIp())
}

func BenchmarkGetInsideIp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetInsideIp())
	}
}

func TestGetMacAddr(t *testing.T) {
	t.Log(GetMacAddr())
}

func BenchmarkGetMacAddr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetMacAddr())
	}
}

func TestIps(t *testing.T) {
	t.Log(Ips())
}

func BenchmarkIps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(Ips())
	}
}
