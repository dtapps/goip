<h1>
<a href="https://www.dtapp.net/">Golang Ip</a>
</h1>

📦 Golang Ip库

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/github.com/dtapps/goip?status.svg)](https://pkg.go.dev/github.com/dtapps/goip)
[![goproxy.cn](https://goproxy.cn/stats/github.com/dtapps/goip/badges/download-count.svg)](https://goproxy.cn/stats/github.com/dtapps/goip)
[![goreportcard.com](https://goreportcard.com/badge/github.com/dtapps/goip	)](https://goreportcard.com/report/github.com/dtapps/goip)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgoip)

#### 安装

```go
go get -v -u github.com/dtapps/goip
```

#### 使用

```go
package main

import (
	"github.com/dtapps/goip"
	"testing"
)

func TestGoIp(t *testing.T) {
	// 获取Mac地址
	t.Log(goip.GetMacAddr())
	// 内网ip
	t.Log(goip.GetInsideIp())
	// 外网ip
	t.Log(goip.GetOutsideIp())
}

```