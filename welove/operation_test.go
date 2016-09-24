package welove

import (
	"testing"
	"io/ioutil"
	"log"
	"github.com/bitly/go-simplejson"
)

func TestTreePost(t *testing.T) {
	res, err := TreePost("562949961343086-2ca7e299a09974dd0", "ac5f34563a4344c4", 2)
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	js, _ := simplejson.NewJson(bytes)
	result, _ := js.Get("result").Int()
	if result != 1 && result != 1001 {
		t.Error("响应值result错误.")
	}
}
