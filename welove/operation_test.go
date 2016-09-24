package welove

import (
	"testing"
	"io/ioutil"
	"log"
	"github.com/bitly/go-simplejson"
	"encoding/json"
)

func TestTreePost(t *testing.T) {
	res, err := TreePost("562949961343086-2ca7e299a09974dd0", "ac5f34563a4344c4", 2)
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	js, _ := simplejson.NewJson(bytes)
	result, _ := js.Get("result").Int()
	if result != 1 && result != 1001 && result != 1002 {
		t.Error("响应值result错误.")
	}
}

func TestHomePost(t *testing.T) {
	res, err := HomePost("562949961343086-2ca7e299a09974dd0", 7, "844424932415867")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	j := make(map[string]interface{})
	json.Unmarshal(body, &j)
	result := j["result"]
	if result != 1201.0 && result != 1.0 {
		log.Fatal("响应值result错误.")
	}
}

func TestRandomHouse(t *testing.T) {
	_, ok := RandomHouse("562949961343086-2ca7e299a09974dd0")
	if !ok {
		t.Error()
	}
}

func TestVisit(t *testing.T) {
	id, ok := RandomHouse("562949961343086-2ca7e299a09974dd0")
	if !ok {
		t.Error()
	}
	res, err := Visit("562949961343086-2ca7e299a09974dd0", id)
	if err != nil {
		t.Error(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	js, err := simplejson.NewJson(bytes)
	if err != nil {
		t.Error(err)
	}
	result, _ := js.Get("result").Int()
	if result != 1 && result != 1201 {
		t.Error("响应值result错误.")
	}
}