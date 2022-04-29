package gorequest

import (
	"net/url"
)

type Headers map[string]string

func NewHeaders() Headers {
	P := make(Headers)
	return P
}

func (p Headers) Set(key, value string) {
	p[key] = value
}

func (p Headers) SetHeaders(headers Headers) {
	for key, value := range headers {
		p[key] = value
	}
}

func (p Headers) GetQuery() string {
	u := url.Values{}
	for k, v := range p {
		u.Set(k, v)
	}
	return u.Encode()
}
