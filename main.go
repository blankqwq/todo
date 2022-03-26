package main

import (
	"log"
	"os"
)

var (
	TodoApp    *Todo
	ConfigData *Config
)

func main() {
	defer func() {
		if TodoApp != nil {
			_ = TodoApp.Done()
		}
	}()
	// 解析配置文件
	ConfigData = &Config{}
	// 初始化数据
	TodoApp = NewTodo(ConfigData)
	cli := NewCli()
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
