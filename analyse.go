package goip

import (
	"strconv"
)

type AnalyseResult struct {
	Ip                string  `json:"ip"`                 // ip
	Continent         string  `json:"continent"`          // 大陆
	Country           string  `json:"country"`            // 国家
	Province          string  `json:"province"`           // 省份
	City              string  `json:"city"`               // 城市
	Isp               string  `json:"isp"`                // 运营商
	LocationTimeZone  string  `json:"location_time_zone"` // 位置时区
	LocationLatitude  float64 `json:"location_latitude"`  // 位置纬度
	LocationLongitude float64 `json:"location_longitude"` // 位置经度
}

func (c *Client) Analyse(ip string) AnalyseResult {
	geoIpInfo, _ := c.QueryGeoIp(ip)
	return AnalyseResult{
		Ip:                ip,
		Continent:         geoIpInfo.Continent.Name,
		Country:           geoIpInfo.Country.Name,
		Province:          geoIpInfo.Province.Name,
		City:              geoIpInfo.City.Name,
		LocationTimeZone:  geoIpInfo.Location.TimeZone,
		LocationLatitude:  geoIpInfo.Location.Latitude,
		LocationLongitude: geoIpInfo.Location.Longitude,
	}
}

// CheckIpv4 检查数据是不是IPV4
func (c *Client) CheckIpv4(ips string) bool {
	if len(ips) > 3 {
		return false
	}
	nums, err := strconv.Atoi(ips)
	if err != nil {
		return false
	}
	if nums < 0 || nums > 255 {
		return false
	}
	if len(ips) > 1 && ips[0] == '0' {
		return false
	}
	return true
}

// CheckIpv6 检测是不是IPV6
func (c *Client) CheckIpv6(ips string) bool {
	if ips == "" {
		return true
	}
	if len(ips) > 4 {
		return false
	}
	for _, val := range ips {
		if !((val >= '0' && val <= '9') || (val >= 'a' && val <= 'f') || (val >= 'A' && val <= 'F')) {
			return false
		}
	}
	return true
}
