package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"strconv"
)

// NewOriginCli 原生解析args
func NewOriginCli(args []string) {
	//pass
}

// NewCli 获取第三方的cli程序
func NewCli() *cli.App {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:  "add",
			Usage: "add a todo",
			Action: func(ctx *cli.Context) error {
				// 调用todo
				err := TodoApp.CreateTodo(ctx.Args().First())
				if err != nil {
					panic(err)
				}
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "get todo list",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:       "all",
					Usage:      "",
					Required:   false,
					HasBeenSet: false,
				},
			},
			Action: func(ctx *cli.Context) error {
				// 调用todo
				var res []string
				if ctx.Bool("all") {
					res, _ = TodoApp.SelectByStatus()
				} else {
					res, _ = TodoApp.SelectByStatus(InCompleteStatus)
				}
				if res != nil {
					for i := range res {
						fmt.Println(res[i])
					}
				} else {
					fmt.Println("没有数据")
				}
				return nil
			},
		},
		{
			Name:  "done",
			Usage: "done a todo",
			Action: func(ctx *cli.Context) error {
				// 调用todo
				idStr := ctx.Args().First()
				id, err := strconv.Atoi(idStr)
				if err != nil {
					log.Fatalf("请输入正确的id")

				}
				err = TodoApp.ChangeTodoStatus(id, CompleteStatus)
				if err != nil {
					log.Fatalf("修改错误: %s", err)
				}
				return nil
			},
		},
		{
			Name:  "login",
			Usage: "login a user",
			Action: func(ctx *cli.Context) error {
				// 调用todo
				// 调用todo
				userName := ctx.Args().First()
				// 请输入密码 todo
				err := TodoApp.Login(userName)
				if err != nil {
					log.Fatalf("登陆错误: %s", err)
				}
				return nil
			},
		},
		{
			Name:  "logout",
			Usage: "logout",
			Action: func(ctx *cli.Context) error {
				// 调用todo
				err := TodoApp.Logout()
				if err != nil {
					log.Fatalf("退出失败: %s", err)
				}
				return nil
			},
		},
	}

	return app
}
