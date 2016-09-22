package welove

import (
	"testing"
)

func TestSig_Encode(t *testing.T) {
	method := "POST"
	u := "http://api.welove520.com/v1/game/house/task"
	content := "access_token=562949961343086-2ca7e299a09974dd0&love_space_id=844424932415867&task_type=5"

	sigEncode := NewSig([]byte("8b5b6eca8a9d1d1f"))
	sig := sigEncode.Encode(method, u, content)
	if sig != "64DO4FLAHIkJADMTuv019C8vAas=" {
		t.Error("Sig_Encode Error.")
	}
}
