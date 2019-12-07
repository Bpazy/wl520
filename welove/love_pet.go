package welove

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type PetTaskResult struct {
	Messages []struct {
		Count      int `json:"count"`
		MsgType    int `json:"msg_type"`
		PetID      int `json:"pet_id"`
		RemainTime int `json:"remain_time"`
		TaskType   int `json:"task_type"`
	} `json:"messages"`
	Result   int    `json:"result"`
	ErrorMsg string `json:"error_msg"`
}

func DoPetTask(accessToken, petId, taskType string) PetTaskResult {
	u := "http://api.welove520.com/v1/game/house/pet/task/do"
	sigEncoder := NewSig([]byte(KEY))
	d1 := Data{"access_token", accessToken}
	d2 := Data{"pet_id", petId}
	d3 := Data{"task_type", taskType}
	sig := sigEncoder.Encode("POST", u, d1, d2, d3)

	data := make(url.Values)
	data.Add("access_token", accessToken)
	data.Add("pet_id", petId)
	data.Add("sig", sig)
	data.Add("task_type", taskType)

	res, err := http.PostForm(u, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	result := PetTaskResult{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		panic(err)
	}
	return result
}

type PetStatus struct {
	Result   int `json:"result"`
	Messages []struct {
		MsgType int `json:"msg_type"`
		Pets    []struct {
			PetID    int `json:"pet_id"`
			PetTasks []struct {
				Count      int `json:"count"`
				TaskType   int `json:"task_type"`
				RemainTime int `json:"remain_time"`
			} `json:"pet_tasks"`
		} `json:"pets,omitempty"`
		Count int `json:"count,omitempty"`
	} `json:"messages"`
}

func GetPetStatus(accessToken string) PetStatus {
	u := "http://api.welove520.com/v1/game/house/pet/task/list"
	sigEncoder := NewSig([]byte(KEY))
	d1 := Data{"access_token", accessToken}
	sig := sigEncoder.Encode("POST", u, d1)

	data := make(url.Values)
	data.Add("access_token", accessToken)
	data.Add("sig", sig)
	res, err := http.PostForm(u, data)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	pet := PetStatus{}
	err = json.Unmarshal(bytes, &pet)
	if err != nil {
		panic(err)
	}
	return pet
}
