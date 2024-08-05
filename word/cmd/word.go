package cmd

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cobra"
	"strings"
	"unicode"
)

var (
	str               string // 传入的字符串
	isUpper           bool   // 转换为大写
	isLower           bool   // 转换为小写
	isUnderline2Camel bool   // 下划线转驼峰
	isCamel2Underline bool   // 驼峰转下划线
)

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().BoolVarP(&isUpper, "upper", "u", false, "是否转换为大写")
	wordCmd.Flags().BoolVarP(&isLower, "lower", "l", false, "是否转换为小写")
	wordCmd.Flags().BoolVarP(&isUnderline2Camel, "u2c", "c", false, "是否下划线转驼峰")
	wordCmd.Flags().BoolVarP(&isCamel2Underline, "c2u", "d", false, "是否驼峰转下划线")
}

var wordCmd = &zdpgo_cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种单词格式转换",
	Run: func(cmd *zdpgo_cobra.Command, args []string) {
		// 这个是运行子命令会直接触发的函数
		if isUpper {
			fmt.Println(strings.ToUpper(str))
		} else if isLower {
			fmt.Println(strings.ToLower(str))
		} else if isUnderline2Camel {
			fmt.Println(underline2Camel(str))
		} else if isCamel2Underline {
			fmt.Println(camel2Underline(str))
		}
	},
}

// 驼峰转下划线
func camel2Underline(s string) any {
	var output []rune
	for i, r := range s {
		if i == 0 {
			// 首字母
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_') // 如果是大写，提前写入一个下划线
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// 下划线转驼峰
func underline2Camel(s string) any {
	s = strings.Replace(s, "_", " ", -1)   // 替换下划线为空格
	s = strings.Title(s)                   // 每个单词首字母大写
	return strings.Replace(s, " ", "", -1) // 移除空格
}
