// Package main is the command entry point for the scaffold service.
package main

// 本文件是编译后二进制的进程入口，只负责装配 CLI 应用并把退出语义交给命令层处理。

import (
	"fmt"
	"os"

	"github.com/rei0721/go-scaffold/pkg/cli"
	"github.com/rei0721/go-scaffold/types/constants"
)

// main 创建 CLI 应用并注册顶层命令，是进程退出码和标准输出错误提示的最后边界。
func main() {
	app := cli.NewApp(constants.AppName)
	app.SetVersion(constants.AppVersion)
	app.SetDescription(constants.AppDescription)

	app.AddCommand(NewAppCommand())
	app.AddCommand(NewDBCommand())

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(cli.GetExitCode(err))
	}
}
