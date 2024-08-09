package main

import (
	"fmt"
	zfile "github.com/zhangdapeng520/zdpgo_file"
	"os"
)

func main() {
	// 判断参数个数
	if len(os.Args) < 2 {
		fmt.Println("请传入要删除的文件后缀")
		return
	}

	// 文件的后缀
	suffix := os.Args[1]

	// 递归删除
	err := zfile.DeleteFileBySuffixRecursion(".", suffix)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("根据后缀递归删除文件成功")

	// 构建二进制文件：go build -o .\bin\dels.exe .\teach\c01_teach_20240808\c02_dels\main.go
	// 使用二进制文件：.\bin\dels.exe .log
}
