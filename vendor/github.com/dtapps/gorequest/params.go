package gorequest

import (
	"encoding/json"
	"github.com/nilorg/sdk/convert"
	"log"
)

type Params map[string]interface{}

func NewParams() Params {
	P := make(Params)
	return P
}

func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

func GetParamsString(src interface{}) string {
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return convert.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
