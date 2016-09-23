package welove

import (
	"testing"
	"crypto/hmac"
	"crypto/sha1"
	"reflect"
)

func TestSig_Encode(t *testing.T) {
	method := "POST"
	u := "http://api.welove520.com/v1/game/house/task"
	d1 := Data{"access_token", "562949961343086-2ca7e299a09974dd0"}
	d2 := Data{"love_space_id", "844424932415867"}
	d3 := Data{"task_type", "5"}

	sigEncode := NewSig([]byte("8b5b6eca8a9d1d1f"))
	sig := sigEncode.Encode(method, u, d1, d2, d3)
	if sig != "64DO4FLAHIkJADMTuv019C8vAas=" {
		t.Error("Sig_Encode2 Error.")
	}
}

func TestNewSig(t *testing.T) {
	key := []byte("8b5b6eca8a9d1d1f")
	mac := hmac.New(sha1.New, key)
	s1 := NewSig(key)
	s2 := &Sig{key, mac}
	if reflect.DeepEqual(s1, s2) {
		t.Error("NewSig Error.")
	}
}