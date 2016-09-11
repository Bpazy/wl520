package main

import "testing"

func TestGetValue(t *testing.T) {
	v := "access_token=token123&sig=sig123&task_type=4&love_space_id=xxaaa22131"
	accessToken, _ := getValue(v, "access_token")
	if accessToken != "token123" {
		t.Error("获取access_token失败")
	}
	loveSpaceId, _ := getValue(v, "love_space_id")
	if loveSpaceId != "xxaaa22131" {
		t.Error("love_space_id")
	}
	taskType, _ := getValue(v, "task_type")
	if taskType != "4" {
		t.Error("task_type")
	}
	sig, _ := getValue(v, "sig")
	if sig != "sig123" {
		t.Error("sig")
	}
}
