package welove

import (
	"net/url"
	"strconv"
	"net/http"
)

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