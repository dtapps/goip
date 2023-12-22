package goip

import (
	"go.dtapp.net/goip/geoip"
)

type ClientConfig struct {
	GeoipAsnPath     string
	GeoipCityPath    string
	GeoipCountryPath string
}

type Client struct {
	geoIpClient *geoip.Client
}

// NewIp 实例化
func NewIp(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	c.geoIpClient, err = geoip.New(config.GeoipAsnPath, config.GeoipCityPath, config.GeoipCountryPath)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close() {
	c.geoIpClient.Close()
}

func (c *Client) GetGeo() *geoip.Client {
	return c.geoIpClient
}
