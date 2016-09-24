package welove

import (
	"strconv"
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

func TreePost(accessToken, appKey string, op int) (*http.Response, error) {
	u := "http://api.welove520.com/v1/game/tree/op"
	sigEncoder := NewSig([]byte("8b5b6eca8a9d1d1f"))
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


func HomePost(accessToken string, taskType int, loveSpaceId string) (*http.Response, error) {
	u := "http://api.welove520.com/v1/game/house/task"
	sigEncoder := NewSig([]byte("8b5b6eca8a9d1d1f"))
	d1:=Data{"access_token", accessToken}
	d2:=Data{"love_space_id", loveSpaceId}
	d3:=Data{"task_type", strconv.Itoa(taskType)}
	sig := sigEncoder.Encode("POST", u, d1,d2,d3)

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
	sigEncoder := NewSig([]byte("8b5b6eca8a9d1d1f"))
	d1 := Data{"access_token", accessToken}
	d2 := Data{"love_space_id", "random"}
	sig := sigEncoder.Encode("POST", u, d1, d2)

	values := make(url.Values)
	values.Add("access_token", accessToken)
	values.Add("love_space_id", "random")
	values.Add("sig", sig)
	res, err := http.PostForm(u, values)
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	js, err := simplejson.NewJson(bytes)
	if err != nil {
		log.Fatal(err)
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
	sigEncoder := NewSig([]byte("8b5b6eca8a9d1d1f"))
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
