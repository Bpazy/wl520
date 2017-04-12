package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"os"
	"github.com/robfig/cron"
	"os/exec"
	"strings"
)

func main() {
	wl520Crons := readConfig("wl520cron.json")
	for _, v := range wl520Crons {
		log.Printf("%s:%s\n", v.Cron, v.Cmd)
		c := cron.New()
		c.AddFunc(v.Cron, func() {
			cmd := exec.Command("welove520", strings.Split(v.Cmd, " ")...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(out))
		})
		c.Start()
	}
	select {}
}

//读取配置文件
func readConfig(configPath string) []Wl520Cron {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		createCronFile(configPath)
		log.Fatalf("配置文件不存在，已创建默认配置文件%s\n", configPath)
	}
	wl520Crons := make([]Wl520Cron, 0)
	json.Unmarshal(bytes, &wl520Crons)
	return wl520Crons
}

type Wl520Cron struct {
	Cron string `json:"cron"`
	Cmd  string `json:"cmd"`
}

//创建默认配置文件
func createCronFile(path string) {
	var f, _ = os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModeAppend)
	defer f.Close()
	wl520Crons := []Wl520Cron{{Cron: "* */30 * * * *", Cmd: "-a -p"},
				  {Cron: "* 0 1,13 * * *", Cmd: "-t -v=20 -farm-sign"}}
	bytes, _ := json.MarshalIndent(wl520Crons, "", "  ")
	f.Write(bytes)
}
