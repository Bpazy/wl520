微爱GoGoGo [![Build Status](https://travis-ci.org/Bpazy/welove520.svg?branch=master)](https://travis-ci.org/Bpazy/welove520)
=============

### 下载并使用
```
go get -u github.com/Bpazy/welove520
```

### 示例
```
1. 生成配置文件 
    welove520  -s -out welove.json
2. 根据配置文件完成任务 
    welove520 -a -c welove.json
```

### 配置文件JSON格式
```
可以通过 welove520  -s -out welove.json 命令生成配置文件
{
  "access_token": "562949961343086-21275eda53f055455f",
  "app_key": "ac5f34563a4344c4",
  "task_type": [
    1,
    4,
    5,
    6,
    7,
    11
  ]
}
```

### 帮助
```
Usage of welove520:
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
  -p    完成宠物任务
  -port string
        我们的家Http代理端口号 (default ":8080")
  -s    启动我们的家HTTP代理
  -t    完成爱情树任务
  -v int
        每日拜访次数 (default -1)
```

微爱API接口 [点击此处](https://github.com/Bpazy/welove520/blob/master/API.md)
