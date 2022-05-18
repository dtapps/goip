package goip

import (
	"net"
	"testing"
)

func lod() *Client {
	return NewIp()
}

func TestIp(t *testing.T) {
	client := lod()
	t.Logf("%+v", net.ParseIP("61.241.55.180").To4())
	t.Logf("%+v", net.ParseIP("240e:3b4:38e4:3295:7093:af6c:e545:f2e9").To16())
	t.Logf("%+v", client.V4db.Find("61.241.55.180"))
	t.Logf("%+v", client.V4db.Find("116.7.98.130"))
	t.Logf("%+v", client.V6db.Find("240e:3b4:38e4:3295:7093:af6c:e545:f2e9"))
	t.Log(client.V4Region.MemorySearch("61.241.55.180"))
}

func BenchmarkIpv4(b *testing.B) {
	client := lod()
	for i := 0; i < b.N; i++ {
		b.Log(client.V4db.Find("61.241.55.180"))
	}
}

func BenchmarkIpv5(b *testing.B) {
	client := lod()
	for i := 0; i < b.N; i++ {
		b.Log(client.V6db.Find("240e:3b4:38e4:3295:7093:af6c:e545:f2e9"))
	}
}
