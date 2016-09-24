微爱GoGoGo [![Build Status](https://travis-ci.org/Bpazy/welove520.svg?branch=master)](https://travis-ci.org/Bpazy/welove520)
=============

### 下载并使用
```
go get github.com/Bpazy/welove520
```

### 示例
```
welove520 -a -c=welove.json
```

### 配置文件JSON格式
```
{
  "access_token": "562949961343086-2ca7e299a09974dd0",
  "love_space_id": "844424932415867",
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
  -alias string
        我们的家生成配置文件详细配置的别名 (default "default")
  -c string
        配置文件位置 (default "welove.json")
  -o string
        日志路径 (default "welove.log")
  -path string
        我们的家生成的配置文件路径 (default "welove.toml")
  -port string
        我们的家Http代理端口号 (default ":8080")
  -s    启动我们的家HTTP代理
  -v int
        每日拜访次数 (default -1)

```

微爱API接口 [点击此处](https://github.com/Bpazy/welove520/blob/master/API.md)
