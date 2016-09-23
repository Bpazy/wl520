package welove

import (
	"hash"
	"crypto/hmac"
	"crypto/sha1"
	"net/url"
	"encoding/base64"
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

func (l *Sig)Encode(method, u string, data ...Data) string {
	var content string
	for _, v := range data {
		content = content + v.key + "=" + v.value + "&"
	}
	content = content[0:len(content) - 1]
	l.myMac.Write([]byte(method + "&" + url.QueryEscape(u) + "&" + url.QueryEscape(content)))
	return base64.StdEncoding.EncodeToString(l.myMac.Sum(nil))
}

type Data struct {
	key   string
	value string
}