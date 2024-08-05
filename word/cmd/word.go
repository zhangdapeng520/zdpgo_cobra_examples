package cmd

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cobra"
	"strings"
)

var (
	str     string // 传入的字符串
	isUpper bool   // 转换为大小
)

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().BoolVarP(&isUpper, "upper", "u", false, "是否转换为大写")
}

var wordCmd = &zdpgo_cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种单词格式转换",
	Run: func(cmd *zdpgo_cobra.Command, args []string) {
		// 这个是运行子命令会直接触发的函数
		if isUpper {
			fmt.Println(strings.ToUpper(str))
		}
	},
}
