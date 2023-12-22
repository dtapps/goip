package geoip

import (
	"github.com/oschwald/geoip2-golang"
)

type Client struct {
	asnDb     *geoip2.Reader
	cityDb    *geoip2.Reader
	countryDb *geoip2.Reader
}

func New(asnFilepath string, cityFilepath string, countryFilepath string) (*Client, error) {

	var err error
	c := &Client{}

	c.asnDb, err = geoip2.Open(asnFilepath)
	if err != nil {
		return nil, err
	}

	c.cityDb, err = geoip2.Open(cityFilepath)
	if err != nil {
		return nil, err
	}

	c.countryDb, err = geoip2.Open(countryFilepath)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (c *Client) Close() {
	c.asnDb.Close()
	c.cityDb.Close()
	c.countryDb.Close()
}
