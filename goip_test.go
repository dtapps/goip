package goip

import (
	"net"
	"testing"
)

func lod() *Client {
	return NewIp()
}

func TestTo4(t *testing.T) {
	t.Logf(string(net.ParseIP("61.241.55.180").To4()))
}

func BenchmarkTo4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Logf(string(net.ParseIP("61.241.55.180").To4()))
	}
}

func TestTo16(t *testing.T) {
	t.Logf(string(net.ParseIP("240e:3b4:38e4:3295:7093:af6c:e545:f2e9").To16()))
}

func BenchmarkTo16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Logf(string(net.ParseIP("240e:3b4:38e4:3295:7093:af6c:e545:f2e9").To16()))
	}
}

func TestIpv4(t *testing.T) {
	client := lod()
	t.Log(client.V4db.Find("116.7.98.130"))
}

func BenchmarkIpv4(b *testing.B) {
	client := lod()
	for i := 0; i < b.N; i++ {
		b.Log(client.V4db.Find("116.7.98.130"))
	}
}

func TestV4Region(t *testing.T) {
	client := lod()
	t.Log(client.V4Region.MemorySearch("61.241.55.180"))
}

func BenchmarkV4Region(b *testing.B) {
	client := lod()
	for i := 0; i < b.N; i++ {
		b.Log(client.V4Region.MemorySearch("61.241.55.180"))
	}
}

func TestIpv6(t *testing.T) {
	client := lod()
	t.Log(client.V6db.Find("116.7.98.130"))
}

func BenchmarkIpv6(b *testing.B) {
	client := lod()
	for i := 0; i < b.N; i++ {
		b.Log(client.V6db.Find("240e:3b4:38e4:3295:7093:af6c:e545:f2e9"))
	}
}
