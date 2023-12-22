package goip

import (
	"go.dtapp.net/goip/geoip"
)

// QueryGeoIp ip2region
// https://www.maxmind.com/
func (c *Client) QueryGeoIp(ipAddress string) (result geoip.QueryCityResult, err error) {
	query, err := c.geoIpClient.QueryCity(ipAddress)
	if err != nil {
		return geoip.QueryCityResult{}, err
	}
	return query, nil
}
