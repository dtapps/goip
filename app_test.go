package goip

import (
	"net"
	"testing"
)

var app App

func TestOnlineDownload(t *testing.T) {
	//t.Log(app.V4db.OnlineDownload())     // 在线下载ipv4数据库
	//t.Log(app.V6db.OnlineDownload())     // 在线下载ipv6数据库
	//t.Log(app.V4Region.OnlineDownload()) // 在线下载ip2region数据库
}

func TestIp(t *testing.T) {
	app.InitIPData()
	t.Logf("%+v", net.ParseIP("61.241.55.180").To4())
	t.Logf("%+v", net.ParseIP("240e:3b4:38e4:3295:7093:af6c:e545:f2e9").To16())
	t.Logf("%+v", app.V4db.Find("61.241.55.180"))
	t.Logf("%+v", app.V4db.Find("116.7.98.130"))
	t.Logf("%+v", app.V6db.Find("240e:3b4:38e4:3295:7093:af6c:e545:f2e9"))
	t.Log(app.V4Region.MemorySearch("61.241.55.180"))
}
