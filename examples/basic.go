package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/willianwong/gtranslate"
)

// 代理服务器
const proxyServer = "http-dyn.abuyun.com:9020"

// 代理隧道验证信息
const proxyUser  = "vas"
const proxyPass  = "casfa"

type AbuyunProxy struct {
	AppID string
	AppSecret string
}

func (p AbuyunProxy) ProxyClient() http.Client {
	proxyUrl, _ := url.Parse("http://"+ p.AppID +":"+ p.AppSecret +"@"+ proxyServer)
	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}

func main() {
	// 初始化 proxy http client
	client := AbuyunProxy{AppID: proxyUser, AppSecret: proxyPass}.ProxyClient()
	text := "Hello World"
	translated, err := gtranslate.TranslateWithFromTo(
		&client,
		text,
		gtranslate.FromTo{
			From: "en",
			To:   "zh-CN",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | zh-CN: %s \n", text, translated)
	// en: Hello World | ja: こんにちは世界
}
