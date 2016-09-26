package welove

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"hash"
	"io"
	"net/url"
	"os"
)

type CustomOutput struct {
	out []io.Writer
}

func New(out ...io.Writer) CustomOutput {
	return CustomOutput{out}
}

func (c *CustomOutput) Add(out io.Writer) {
	c.out = append(c.out, out)
}

func (c *CustomOutput) Write(p []byte) (int, error) {
	var n int = 0
	var err error = nil
	for _, v := range c.out {
		n, err = v.Write(p)
	}
	return n, err
}

func DefaultLog(path string) CustomOutput {
	var file, _ = os.OpenFile(path, os.O_APPEND | os.O_CREATE | os.O_RDWR, os.ModeAppend)
	return New(os.Stdout, file)
}

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
	content = content[0 : len(content) - 1]
	l.myMac.Write([]byte(method + "&" + url.QueryEscape(u) + "&" + url.QueryEscape(content)))
	return base64.StdEncoding.EncodeToString(l.myMac.Sum(nil))
}

type Data struct {
	key   string
	value string
}
