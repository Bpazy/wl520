微爱GoGoGo
=============
[![Build Status](https://github.com/Bpazy/wl520/workflows/Build/badge.svg)](https://github.com/Bpazy/wl520/actions)

### 下载并使用
前往[Release](https://github.com/Bpazy/wl520/releases)下载适合您系统的版本。

### 示例
```
1. 生成配置文件 
    wl520  -s -out welove.json
2. 根据配置文件完成任务 
    wl520 -a -c welove.json
```

### 配置文件JSON格式
```
可以通过 wl520  -s -out welove.json 命令生成配置文件
{
  "access_token": "562949961343086-****4d8f8d23b747a3",
  "app_key": "ac5f345****344c4",
   // 阿里大于配置文件，如不需要短信提示功能，可不填
  "mobile": "18*****2070",
  "lover_mobile": "15*****8066",
  "dayu_app_key": "24****26",
  "dayu_app_secret": "817711************8711d3f5228",
  "dayu_sunlight_template_code": "SMS_********",
  "dayu_water_template_code": "SMS_********",
  "dayu_home_template_code": "SMS_********"
}
```

### 帮助
```
Usage of wl520:
  -a    完成所有我们的家互动任务
  -buy int
        农场购买物品ID
  -c string
        配置文件位置 (default "welove.json")
  -coin int
        农场被购买物品ID的价格上限(闭区间) (default -1)
  -log string
        日志路径 (default "welove.log")
  -out string
        生成的配置文件路径 (default "welove.json")
  -p    完成宠物任务（需要完成的时候才会消耗物品）
  -port string
        我们的家Http代理端口号 (default ":8080")
  -s    启动我们的家HTTP代理
  -t    完成爱情树任务
  -v int
        每日拜访次数 (default -1)
  -farm-sign
        农场签到
```

### 添加定时任务
1\. 你可以使用`Linux`的工具`cron`:
```
#每30分钟检测完成我们的家所有任务和宠物任务
*/30 * * * * /usr/bin/wl520 -c /etc/welove.json -log /home/han/welove/welove_han.log -a -p

#每天凌晨1点和下午13点(两次是为了防止请求失败)检测并完成爱情树任务, 拜访20次任务, 农场签到任务
0 1,13 * * * /usr/bin/wl520 -c /etc/welove.json -log /home/han/welove/welove_han.log -t -v=20 -farm-sign
```
2\. 或者使用本项目提供的`wl520cron`:    
&emsp;2.1 获取`wl520cron`: `go get github.com/Bpazy/wl520/wl520cron`;    
&emsp;2.2 设置`wl520cron`的配置文件`wl520cron.json`;    
&emsp;2.3 运行`wl520cron`, `wl520`必须在`PATH`下。

`wl520cron.json`格式:
```
 [
   {
     "cron": "* */30 * * * *",
     "cmd": "-a -p"                // cmd为wl520的命令
   },
   {
     "cron": "* 0 1,13 * * *",
     "cmd": "-t -v=20 -farm-sign"
   }
 ]
 ```
`cron`表达式说明:
```
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
```

### Docker
首先创建配置文件`wl520.json`和`wl520cron.json`并存放在一个目录中，然后运行`docker`:
```
docker run -d -v /your-directory/wl520:/etc/wl520 bpazy/wl520
```

### 编译
```shell
$ go build github.com/Bpazy/wl520/cmd/wl520
```
