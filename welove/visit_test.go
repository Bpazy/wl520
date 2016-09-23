package welove

import (
	"testing"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

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