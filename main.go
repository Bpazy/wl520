package main

import (
	"encoding/json"
	"flag"
	"github.com/Bpazy/welove520/welove"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	isServer := flag.Bool("s", false, "启动我们的家HTTP代理")
	path := flag.String("out", "welove.json", "生成的配置文件路径")
	port := flag.String("port", ":8080", "我们的家Http代理端口号")
	allTask := flag.Bool("a", false, "完成所有我们的家互动任务")
	configPath := flag.String("c", "welove.json", "配置文件位置")
	visitTimes := flag.Int("v", -1, "每日拜访次数")
	outputPath := flag.String("log", "welove.log", "日志路径")
	tree := flag.Bool("t", false, "完成爱情树任务")
	pet := flag.Bool("p", false, "完成宠物任务")
	flag.Parse()

	//是否开启代理服务器
	if *isServer {
		welove.ServerRun(*path, *port)
	}

	love := initConfig(*outputPath, *configPath)
	//完成互动任务
	if *allTask {
		doAllTasks(love)
	}
	//拜访任务
	if *visitTimes != -1 {
		doVisit(*visitTimes, love)
	}
	//爱情树任务
	if *tree {
		doTreePost(love)
	}
	//宠物任务
	if *pet {
		doPetTasks(love)
	}
}

func doTreePost(love welove.Love) {
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
func initConfig(outputPath, configPath string) welove.Love {
	//配置日志
	output := welove.DefaultLog(outputPath)
	log.SetOutput(&output)

	//读取配置文件
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	love := welove.Love{}
	json.Unmarshal(bytes, &love)
	return love
}

func doVisit(visitTimes int, love welove.Love) {
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

func doAllTasks(love welove.Love) {
	res, err := welove.GetLoveSpaceIdRaw(love.AccessToken, love.AppKey)
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	loveSpaceId := welove.GetLoveSpaceId(string(bytes))
	for _, v := range love.TaskType {
		res, err := welove.HomePost(love.AccessToken, v, loveSpaceId)
		if err != nil {
			log.Printf("任务%d错误\n", love.TaskType)
		}
		bytes, _ := ioutil.ReadAll(res.Body)
		m := make(map[string]interface{})
		json.Unmarshal(bytes, &m)
		log.Printf("任务%d, Raw: %s\n", v, string(bytes))
	}
}

func doPetTasks(love welove.Love) {
	petStatus := welove.GetPetStatus(love.AccessToken)
	log.Printf("宠物状态Raw: %+v\n", petStatus)
	pet := petStatus.Messages[0].Pets[0]
	for _, v := range pet.PetTasks {
		if v.RemainTime != 0 {
			continue
		}
		taskResult := welove.DoPetTask(love.AccessToken, strconv.Itoa(pet.PetID), strconv.Itoa(v.TaskType))
		log.Printf("宠物任务%d, Raw: %+v\n", v.TaskType, taskResult)
	}
}
