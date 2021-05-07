package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type IpSvc struct {
	Name   string
	local  string
	global func(ip string) string
}

func (svc IpSvc) LocalInfo() string {
	return ipGet(svc.local)
}

func (svc IpSvc) GlobalInfo(ip string) string {
	return ipGet(svc.global(ip))
}

var IpipSvc = IpSvc{
	Name:  "ipip.net",
	local: (func() string {
		rand.Seed(time.Now().UnixNano())
		return fmt.Sprintf("http://myip.ipip.net/?_=%d", rand.Int())
	})(),
	global: func(ip string) string {
		return fmt.Sprintf("http://freeapi.ipip.net/%s", ip)
	},
}
var IpInfoSvc = IpSvc{
	Name:  "ipinfo.io",
	local: "http://ipinfo.io",
	global: func(ip string) string {
		return fmt.Sprintf("http://ipinfo.io/%s", ip)
	},
}
var IpFmSvc = IpSvc{
	Name:  "ip.fm",
	local: "https://ip.fm",
	global: func(ip string) string {
		return fmt.Sprintf("http://ip.fm/?ip=%s", ip)
	},
}

//http get 获取 ip 信息
func ipGet(url string) string {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "curl/7.76.1")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	return strings.TrimSpace(string(body))
}

func SvcIpip(ip string) string {
	if "" == ip {
		return IpipSvc.LocalInfo()
	} else {
		return IpipSvc.GlobalInfo(ip)
	}
}

func SvcIpInfo(ip string) string {
	if "" == ip {
		return IpInfoSvc.LocalInfo()
	} else {
		return IpInfoSvc.GlobalInfo(ip)
	}
}

func SvcIpFm(ip string) string {
	if "" == ip {
		return IpFmSvc.LocalInfo()
	} else {
		return IpFmSvc.GlobalInfo(ip)
	}
}
