package main

import (
	"flag"
	"github.com/Bpazy/welove520/welove"
)

func main() {
	isServer := flag.Bool("s", false, "启动我们的家HTTP代理")
	path := flag.String("path", "welove.toml", "我们的家生成的配置文件路径")
	port := flag.String("port", ":8080", "我们的家Http代理端口号")
	alias := flag.String("alias", "default", "我们的家生成配置文件详细配置的别名")
	flag.Parse()
	if (*isServer) {
		welove.ServerRun(path, port, alias)
	}
}
