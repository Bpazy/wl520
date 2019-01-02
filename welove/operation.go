package welove

import (
	bb "bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/elazarl/goproxy.v1"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const KEY = "8b5b6eca8a9d1d1f"

var sChan = make(chan string)

func ServerRun(path, port string, server bool) {
	if !server {
		return
	}
	log.Printf("请将手机Http代理设置为[本机IP%s]\n", port)
	go contentHandler(path)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	proxy.OnRequest().DoFunc(httpHandler)
	log.Fatal(http.ListenAndServe(port, proxy))
}

// 配置文件
type Love struct {
	AccessToken              string `json:"access_token"`
	AppKey                   string `json:"app_key"`
	Mobile                   string `json:"mobile"`
	LoverMobile              string `json:"lover_mobile"`
	DayuAppKey               string `json:"dayu_app_key"`
	DayuAppSecret            string `json:"dayu_app_secret"`
	DayuSunlightTemplateCode string `json:"dayu_sunlight_template_code"`
	DayuWaterTemplateCode    string `json:"dayu_water_template_code"`
	DayuHomeTemplateCode     string `json:"dayu_home_template_code"`
}

func contentHandler(path string) {
	var f, _ = os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModeAppend)
	defer f.Close()
	for v := range sChan {
		accessToken, _ := extractValue(v, "access_token")
		appKey, _ := extractValue(v, "app_key")
		love := Love{}
		love.AccessToken = accessToken
		love.AppKey = appKey
		bytes, _ := json.MarshalIndent(love, "", "  ")
		f.Write(bytes)
		fmt.Println("生成配置文件完毕：" + path)
		os.Exit(0)
	}
}

func extractValue(content, key string) (string, error) {
	r := "(?:" + key + ")=(.+?)(&|$)"
	reg, err := regexp.Compile(r)
	return reg.FindAllStringSubmatch(content, -1)[0][1], err
}

func httpHandler(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	if r.Method != "POST" {
		return r, nil
	}
	if r.Host != "api.welove520.com" {
		return r, nil
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	ori := ioutil.NopCloser(bb.NewBuffer(buf))
	r.Body = ori
	bytesString := bb.Buffer{}
	bytesString.Write(buf)
	content := bytesString.String()
	if strings.Contains(content, "access_token") && strings.Contains(content, "app_key") {
		dContent, err := url.QueryUnescape(content)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Decode [%s] to [%s]\n", content, dContent)
		sChan <- dContent
	}
	return r, nil
}
