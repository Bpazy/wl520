
### 下载并使用
```
go get github.com/Bpazy/welove520
welove520 -path=welove.conf -port=:8080
将手机http代理设置为本机的8080端口
```
微爱请求关键信息在执行后保存在根目录welove.conf中

由Python发送post请求 [Example](https://github.com/Bpazy/welove520_API/blob/master/example/post.py)

添加定时任务每`30`分钟执行一次

`echo "*/30 * * * * python welove.py" >> /var/spool/cron/root`

微爱API接口 [点击此处](https://github.com/Bpazy/welove520_API/blob/master/example/API.md)
