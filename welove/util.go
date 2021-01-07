package welove

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"hash"
	"net/http"
	"net/url"
	"strings"
)

type Sig struct {
	key   []byte
	myMac hash.Hash
}

func NewSig(key []byte) *Sig {
	mac := hmac.New(sha1.New, key)
	love := new(Sig)
	love.myMac = mac
	return love
}

func (l *Sig) Encode(method, u string, data ...Data) string {
	var content string
	for _, v := range data {
		content = content + v.key + "=" + v.value + "&"
	}
	content = content[0 : len(content)-1]
	l.myMac.Write([]byte(method + "&" + url.QueryEscape(u) + "&" + url.QueryEscape(content)))
	return base64.StdEncoding.EncodeToString(l.myMac.Sum(nil))
}

type Data struct {
	key   string
	value string
}

type WlHttpClient struct {
	Client *http.Client
}

func NewWlHttpClient() *WlHttpClient {
	client := &http.Client{}
	wlClient := WlHttpClient{Client: client}
	return &wlClient
}

func (client *WlHttpClient) Post(url string, data url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Welove-UA", "[Device:ONEPLUSA5010][OSV:7.1.1][CV:Android4.0.3][WWAN:0][zh_CN][platform:tencent][WSP:2]")
	return client.Client.Do(req)
}
