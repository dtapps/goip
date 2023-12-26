package goip

import (
	"go.dtapp.net/goip/geoip"
	"go.dtapp.net/goip/qqwry"
)

func (c *Client) GetGeo() *geoip.Client {
	return c.geoIpClient
}

func (c *Client) GetQqWry() *qqwry.Client {
	return c.qqwryClient
}
