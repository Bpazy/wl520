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
	AppKey      string `json:"app_key"`
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
	outputPath := flag.String("o", "welove.log", "日志路径")
	tree := flag.Bool("t", false, "是否完成爱情树任务")
	flag.Parse()

	love := initConfig(*outputPath, *configPath)
	//是否开启代理服务器
	if *isServer {
		welove.ServerRun(*path, *port, *alias)
	}
	//完成互动任务
	if *allTask {
		doAllTask(love)
	}
	//拜访任务
	if *visitTimes != -1 {
		doVisit(*visitTimes, love)
	}
	//爱情树任务
	if *tree {
		doTreePost(love)
	}
}

func doTreePost(love Love) {
	op := []int{1, 2}
	for _, v := range op {
		res, err := welove.TreePost(love.AccessToken, love.AppKey, v)
		if err != nil {
			log.Fatal(err)
		}
		bytes, _ := ioutil.ReadAll(res.Body)
		js, _ := simplejson.NewJson(bytes)
		result, _ := js.Get("result").Int()
		log.Printf("爱情树result: %d, Raw: %s\n", result, string(bytes))
	}
}
func initConfig(outputPath, configPath string) Love {
	//配置日志
	output := welove.DefaultLog(outputPath)
	log.SetOutput(&output)

	//读取配置文件
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	love := Love{}
	json.Unmarshal(bytes, &love)
	return love
}

func doVisit(visitTimes int, love Love) {
	for i := 0; i < visitTimes; i++ {
		if id, ok := welove.RandomHouse(love.AccessToken); ok {
			res, err := welove.Visit(love.AccessToken, id)
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