package welove

import (
	"net/url"
	"strconv"
	"net/http"
)

var u = "http://api.welove520.com/v1/game/house/task"

func HomePost(accessToken string, taskType int, loveSpaceId string) (*http.Response, error) {
	sigEncoder := NewSig("8b5b6eca8a9d1d1f")
	content := accessToken + "&" + strconv.Itoa(taskType) + "&" + loveSpaceId
	sig := sigEncoder.Encode("POST", u, content)

	data := make(url.Values)
	data.Add("access_token", accessToken)
	data.Add("task_type", taskType)
	data.Add("love_space_id", loveSpaceId)
	data.Add("sig", sig)
	res, err := http.PostForm(u, data)
	return res, err
}