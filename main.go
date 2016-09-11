package main

import (
	"bufio"
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"flag"
)

type WeloveValue struct {
	AccessToken string `toml:"access_token"`
	TaskType    string `toml:"task_type"`
	LoveSpaceId string `toml:"love_space_id"`
	Sig         string `toml:"sig"`
}

var sChan = make(chan string)

func contentHandler(path *string) {
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
			m["home"] = value
		case "4":
			m["eat"] = value
		case "5":
			m["sleep"] = value
		case "6":
			m["bath"] = value
		case "7":
			m["rest"] = value
		case "11":
			m["mua"] = value
		default:
			m["unknown, please report an issue"] = value
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
		log.Println(content)
		sChan <- content
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
	path := flag.String("path", "welove.toml", "log文件路径，默认当前路径")
	port := flag.String("port", ":8080", "Http代理端口号")
	flag.Parse()
	log.Printf("请将手机Http代理设置为[本机IP%s]\n", *port)
	go input()
	go contentHandler(path)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	proxy.OnRequest().DoFunc(httpHandler)
	log.Fatal(http.ListenAndServe(*port, proxy))
}
