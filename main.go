package main

import (
	"flag"
	"github.com/Bpazy/welove520/welove"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

type Love struct {
	AccessToken string `json:"access_token"`
	LoveSpaceId string `json:"love_space_id"`
	TaskType    []int `json:"task_type"`
}

func main() {
	isServer := flag.Bool("s", false, "启动我们的家HTTP代理")
	path := flag.String("path", "welove.toml", "我们的家生成的配置文件路径")
	port := flag.String("port", ":8080", "我们的家Http代理端口号")
	alias := flag.String("alias", "default", "我们的家生成配置文件详细配置的别名")
	allTask := flag.Bool("a", false, "完成所有我们的家互动任务")
	configPath := flag.String("c", "welove.json", "配置文件位置")
	flag.Parse()

	if *isServer {
		welove.ServerRun(*path, *port, *alias)
	}
	if *allTask {
		doAllTask(*configPath)
	}
}

func doAllTask(configPath string) {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	love := Love{}
	json.Unmarshal(bytes, &love)
	for _, v := range love.TaskType {
		res, err := welove.HomePost(love.AccessToken, v, love.LoveSpaceId)
		if err != nil {
			fmt.Printf("任务%d错误\n", love.TaskType)
		}
		bytes, _ = ioutil.ReadAll(res.Body)
		m := make(map[string]interface{})
		json.Unmarshal(bytes, &m)
		fmt.Printf("任务%d, Raw: %s\n", v, string(bytes))
	}
}