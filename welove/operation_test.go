package welove

import (
	"testing"
)

func TestGetLoveSpaceId(t *testing.T) {
	body := `{"result":1,"love_space_id":844424932415867,"emotion_last":8,"emotion_cur":0,"user_id":562949961343055,"set_time":1474348805820}`
	id := GetLoveSpaceId(body)
	if id != "844424932415867" {
		t.Error()
	}
}
