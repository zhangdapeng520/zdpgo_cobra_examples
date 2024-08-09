package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义变量
	var name string

	// 从命令行进行参数解析
	flag.StringVar(&name, "name", "默认值", "使用的说明")
	flag.StringVar(&name, "n", "张大鹏", "用户名")

	// 进行解析
	flag.Parse()

	// 查看解析结果
	// go run main.go -n 张大鹏333
	fmt.Println("name=", name)
}
