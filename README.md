[![Build Status](https://travis-ci.org/Bpazy/welove520.svg?branch=master)](https://travis-ci.org/Bpazy/welove520)
### 下载并使用
```
go get github.com/Bpazy/welove520
welove520 -path=welove.toml -port=:8080

将手机http代理设置为本机的8080端口
输入#回车终止程序
```
微爱请求关键信息在执行后默认保存在根目录welove.toml中

由Python发送post请求 [Example](https://github.com/Bpazy/welove520_API/blob/master/example/post.py)

添加定时任务每`30`分钟执行一次

`echo "*/30 * * * * python welove.py" >> /var/spool/cron/root`

微爱API接口 [点击此处](https://github.com/Bpazy/welove520_API/blob/master/example/API.md)
