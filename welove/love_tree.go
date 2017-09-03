package welove

import (
	"net/http"
	"strconv"
	"net/url"
)

func TreePost(accessToken, appKey string, op int) (*http.Response, error) {
	u := "http://api.welove520.com/v1/game/tree/op"
	sigEncoder := NewSig([]byte(KEY))
	d1 := Data{"access_token", accessToken}
	d2 := Data{"app_key", appKey}
	d3 := Data{"op", strconv.Itoa(op)}
	sig := sigEncoder.Encode("POST", u, d1, d2, d3)

	data := make(url.Values)
	data.Add("access_token", accessToken)
	data.Add("app_key", appKey)
	data.Add("op", strconv.Itoa(op))
	data.Add("sig", sig)
	res, err := http.PostForm(u, data)
	return res, err
}
