package ip2region_v2

import (
	"io/ioutil"
	"net/http"
)

func OnlineDownload() {
	resp, err := http.Get("https://ghproxy.com/?q=https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.xdb?raw=true")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = ioutil.WriteFile("./ip2region.xdb", body, 0644)
	if err != nil {
		panic(err)
	}
}
