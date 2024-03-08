<h1>
<a href="https://www.dtapp.net/">Golang</a>
</h1>

📦 Golang IP数据库

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/go.dtapp.net/goip?status.svg)](https://pkg.go.dev/go.dtapp.net/goip)
[![goproxy.cn](https://goproxy.cn/stats/go.dtapp.net/goip/badges/download-count.svg)](https://goproxy.cn/stats/go.dtapp.net/goip)
[![goreportcard.com](https://goreportcard.com/badge/go.dtapp.net/goip	)](https://goreportcard.com/report/go.dtapp.net/goip)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/go.dtapp.net%2Fgoip)

#### 安装

```shell
go get -v -u go.dtapp.net/goip@v1.0.47
```

#### 使用

```go
package main

import (
	"context"
	"go.dtapp.net/goip"
	"testing"
)

func TestGoIp(t *testing.T) {
	// 获取Mac地址
	t.Log(goip.GetMacAddr(context.Background()))
	// 内网ip
	t.Log(goip.GetInsideIp(context.Background()))
	// 外网ip
	t.Log(goip.GetOutsideIp(context.Background()))
}

```