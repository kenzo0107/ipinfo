package ipinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func Get(ip string) IPInfo {
	url := fmt.Sprintf("http://ipinfo.io/%s/json", ip)
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ipinfo IPInfo
	err = json.Unmarshal([]byte(string(b)), &ipinfo)
	if err != nil {
		log.Fatal(err)
	}

	return ipinfo
}
