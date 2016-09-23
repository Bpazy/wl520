package main

import (
	"flag"
	"github.com/Bpazy/welove520/welove"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/bitly/go-simplejson"
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
	visitTimes := flag.Int("v", -1, "每日拜访次数")

	flag.Parse()

	/**
	是否开启代理服务器
	 */
	if *isServer {
		welove.ServerRun(*path, *port, *alias)
	}

	/**
	读取配置文件
	 */
	bytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	love := Love{}
	json.Unmarshal(bytes, &love)

	/**
	完成互动任务
	 */
	if *allTask {
		doAllTask(love)
	}

	/**
	拜访任务
	 */
	if *visitTimes != -1 {
		doVisit(*visitTimes, love.AccessToken)
	}
}

func doVisit(visitTimes int, accessToken string) {
	for i := 0; i < visitTimes; i++ {
		if id, ok := welove.RandomHouse(accessToken); ok {
			res, err := welove.Visit(accessToken, id)
			if err != nil {
				log.Fatal(err)
			}
			bytes, _ := ioutil.ReadAll(res.Body)
			js, _ := simplejson.NewJson(bytes)
			result, _ := js.Get("result").Int()
			log.Printf("拜访result: %d, Raw: %s\n", result, string(bytes))
		}
	}
}

func doAllTask(love Love) {
	for _, v := range love.TaskType {
		res, err := welove.HomePost(love.AccessToken, v, love.LoveSpaceId)
		if err != nil {
			log.Printf("任务%d错误\n", love.TaskType)
		}
		bytes, _ := ioutil.ReadAll(res.Body)
		m := make(map[string]interface{})
		json.Unmarshal(bytes, &m)
		log.Printf("任务%d, Raw: %s\n", v, string(bytes))
	}
}