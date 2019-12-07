package welove

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func HomePost(accessToken string, taskType int, loveSpaceId string) (*http.Response, error) {
	u := "http://api.welove520.com/v1/game/house/task"
	sigEncoder := NewSig([]byte(KEY))
	d1 := Data{"access_token", accessToken}
	d2 := Data{"love_space_id", loveSpaceId}
	d3 := Data{"task_type", strconv.Itoa(taskType)}
	sig := sigEncoder.Encode("POST", u, d1, d2, d3)

	data := make(url.Values)
	data.Add("access_token", accessToken)
	data.Add("task_type", strconv.Itoa(taskType))
	data.Add("love_space_id", loveSpaceId)
	data.Add("sig", sig)
	res, err := http.PostForm(u, data)
	return res, err
}

func RandomHouse(accessToken string) (string, bool) {
	var u = "http://api.welove520.com/v1/game/house/info"
	sigEncoder := NewSig([]byte(KEY))
	d1 := Data{"access_token", accessToken}
	d2 := Data{"love_space_id", "random"}
	sig := sigEncoder.Encode("POST", u, d1, d2)

	values := make(url.Values)
	values.Add("access_token", accessToken)
	values.Add("love_space_id", "random")
	values.Add("sig", sig)
	res, err := http.PostForm(u, values)
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	js, err := simplejson.NewJson(bytes)
	if err != nil {
		panic(err)
	}
	arr, err := js.Get("messages").Array()
	house, ok := arr[0].(map[string]interface{})["house"].(map[string]interface{})
	if ok {
		id, ok := house["love_space_id"].(string)
		return id, ok
	} else {
		return "", ok
	}
}

func Visit(accessToken, loveSpaceId string) (*http.Response, error) {
	u := "http://api.welove520.com/v1/game/house/task"

	d1 := Data{"task_type", "8"}
	d2 := Data{"house_num", "0"}
	d3 := Data{"access_token", accessToken}
	d4 := Data{"love_space_id", loveSpaceId}
	sigEncoder := NewSig([]byte(KEY))
	sig := sigEncoder.Encode("POST", u, d3, d2, d4, d1)

	values := make(url.Values)
	values.Add("task_type", "8")
	values.Add("house_num", "0")
	values.Add("access_token", accessToken)
	values.Add("love_space_id", loveSpaceId)
	values.Add("sig", sig)
	res, err := http.PostForm(u, values)
	return res, err
}

// 我们的家
func GetLoveSpaceIdRaw(accessToken, appKey string) (*http.Response, error) {
	u := "http://api.welove520.com/v5/useremotion/getone"
	d1 := Data{"access_token", accessToken}
	d2 := Data{"app_key", appKey}
	d3 := Data{"user_id", "0"}
	sigEncoder := NewSig([]byte(KEY))
	sig := sigEncoder.Encode("POST", u, d1, d2, d3)

	data := make(url.Values)
	data.Add("access_token", accessToken)
	data.Add("app_key", appKey)
	data.Add("user_id", "0")
	data.Add("sig", sig)
	res, err := http.PostForm(u, data)
	return res, err
}

func GetLoveSpaceId(body string) string {
	js, _ := simplejson.NewJson([]byte(body))
	loveSpaceId, _ := js.Get("love_space_id").Float64()
	return strconv.Itoa(int(loveSpaceId))
}
