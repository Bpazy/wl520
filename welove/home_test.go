package welove

import (
	"testing"
	"log"
	"io/ioutil"
	"encoding/json"
)

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
