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
	buyItemId := flag.Int("buy", 0, "农场购买物品ID")
	coin := flag.Int("coin", -1, "农场被购买物品ID的价格上限(闭区间)")
	doFarmSign := flag.Bool("farm-sign", false, "农场签到")
	flag.Parse()

	welove.ServerRun(*path, *port, *isServer)         //是否开启代理服务器
	love := initConfig(*outputPath, *configPath)      //读取配置文件

	go buyItem(love, *buyItemId, *coin, *buyItemId)   //购买指定物品
	go doAllTasks(love, *allTask)                     //完成互动任务
	go doVisit(*visitTimes, love)                     //拜访任务
	go doTreePost(love, *tree)                        //爱情树任务
	go doPetTasks(love, *pet)                         //宠物任务
	go farmSign(love, *doFarmSign)                    //农场签到
}

func farmSign(love welove.Love, do bool) {
	if !do {
		return
	}
	res, err := welove.FarmSign(love.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	log.Printf("农场签到: %s\n", string(bytes))
}

func buyItem(love welove.Love, itemId, coin, buyItemId int) {
	if buyItemId == 0 {
		return
	}
	items := welove.QueryItems(love.AccessToken).Messages[0].AdItems
	for _, v := range items {
		if v.ItemID == itemId && v.Coin <= coin {
			status := welove.BuyItem(love.AccessToken, v.SellerFarmID, v.ID)
			log.Printf("农场购买result: %d, Raw: %+v\n", status.Result, status)
		}
	}
}

func doTreePost(love welove.Love, tree bool) {
	if !tree {
		return
	}
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
	if visitTimes == -1 {
		return
	}
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

func doAllTasks(love welove.Love, allTask bool) {
	if !allTask {
		return
	}
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

func doPetTasks(love welove.Love, doPet bool) {
	if !doPet {
		return
	}
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
