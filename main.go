package main

import (
	"bufio"
	"bytes"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"net/url"
)

type WeloveValue struct {
	AccessToken string `toml:"access_token"`
	TaskType    string `toml:"task_type"`
	LoveSpaceId string `toml:"love_space_id"`
	Sig         string `toml:"sig"`
}

var sChan = make(chan string)

func contentHandler(path, alias *string) {
	var f, _ = os.Create(*path)
	defer f.Close()
	var encoder = toml.NewEncoder(f)
	for v := range sChan {
		var m = make(map[string]WeloveValue)
		accessToken, _ := getValue(v, "access_token")
		loveSpaceId, _ := getValue(v, "love_space_id")
		taskType, _ := getValue(v, "task_type")
		sig, _ := getValue(v, "sig")
		value := WeloveValue{accessToken, taskType, loveSpaceId, sig}
		switch taskType {
		case "1":
			m[*alias + ":home"] = value
		case "4":
			m[*alias + ":eat"] = value
		case "5":
			m[*alias + ":sleep"] = value
		case "6":
			m[*alias + ":bath"] = value
		case "7":
			m[*alias + ":rest"] = value
		case "11":
			m[*alias + ":mua"] = value
		default:
			m[*alias + ":unknown, please report an issue"] = value
		}
		encoder.Encode(m)
	}
	os.Exit(0)
}

func getValue(content, key string) (string, error) {
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
	ori := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = ori
	bytesString := bytes.Buffer{}
	bytesString.Write(buf)
	content := bytesString.String()
	if strings.Contains(content, "sig") &&
		strings.Contains(content, "love_space_id") &&
		strings.Contains(content, "access_token") &&
		strings.Contains(content, "task_type") {
		dContent, err := url.QueryUnescape(content)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Decode [%s] to [%s]\n", content, dContent)
		sChan <- dContent
	}
	return r, nil
}

func input() {
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		s := string(data)
		if s == "#" {
			close(sChan)
		}
	}
}

func main() {
	path := flag.String("path", "welove.toml", "生成的配置文件路径")
	port := flag.String("port", ":8080", "Http代理端口号")
	alias := flag.String("alias", "default", "生成配置文件详细配置的别名")
	flag.Parse()
	log.Printf("请将手机Http代理设置为[本机IP%s]\n", *port)
	go input()
	go contentHandler(path, alias)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	proxy.OnRequest().DoFunc(httpHandler)
	log.Fatal(http.ListenAndServe(*port, proxy))
}
