package main

import (
	"fmt"
	"github.com/hnmaonanbei/go-youdao-fanyi"
)

func main() {
	opts := ydfanyi.NewOptions("", "", "")
	// opts.proxy = ""
	//opts.From = ydfanyi.EN
	//opts.To = ydfanyi.ZH
	res, _ := ydfanyi.Do("hello ", opts)
	// 翻译结果
	fmt.Println(res.String())
	// 详细信息
	fmt.Println(res.SmartResult.Entries)
	// out

	//	你好
	//	[ int. 喂；哈罗，你好，您好（表示问候， 惊奇或唤起注意时的用语）
	//	n. （Hello）（法）埃洛（人名）
	//]

}
