package main

import (
	"fmt"
	"github.com/hnmaonanbei/go-youdao-fanyi"
)

func main() {
	opts := ydfanyi.NewOptions("", "", "")

	// 设置代理(可选)
	// opts.proxy = ""

	// 自动检测语言
	res, _ := ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 英语
	opts.To = ydfanyi.EN
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 日语
	opts.To = ydfanyi.JA
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 韩语
	opts.To = ydfanyi.KO
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 法语
	opts.To = ydfanyi.FR
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 德语
	opts.To = ydfanyi.DE
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 俄语
	opts.To = ydfanyi.RU
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 西班牙语
	opts.To = ydfanyi.ES
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 葡萄牙语
	opts.To = ydfanyi.PT
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 意大利语
	opts.To = ydfanyi.IT
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 越南语
	opts.To = ydfanyi.VI
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 印尼语
	opts.To = ydfanyi.ID
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 阿拉伯语
	opts.To = ydfanyi.AR
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	// 中文 -> 荷兰语
	opts.To = ydfanyi.NL
	res, _ = ydfanyi.Do("你好世界", opts)
	fmt.Println(res.String())

	//
	opts.To = ydfanyi.ZH

	// 英语 -> 中文
	opts.From = ydfanyi.EN
	res, _ = ydfanyi.Do("Hello World", opts)
	fmt.Println(res.String())

	// 日语 -> 中文
	opts.From = ydfanyi.JA
	res, _ = ydfanyi.Do("こんにちは世界", opts)
	fmt.Println(res.String())

	// 韩语 -> 中文
	opts.From = ydfanyi.KO
	res, _ = ydfanyi.Do("하이 월드", opts)
	fmt.Println(res.String())

	// 法语 -> 中文
	opts.From = ydfanyi.FR
	res, _ = ydfanyi.Do("Hello world", opts)
	fmt.Println(res.String())

	// 德语 -> 中文
	opts.From = ydfanyi.DE
	res, _ = ydfanyi.Do("Hallo, welt.", opts)
	fmt.Println(res.String())

	// 俄语 -> 中文
	opts.From = ydfanyi.RU
	res, _ = ydfanyi.Do("Привет, мир.", opts)
	fmt.Println(res.String())

	// 西班牙语 -> 中文
	opts.From = ydfanyi.ES
	res, _ = ydfanyi.Do("Hola mundo", opts)
	fmt.Println(res.String())

	// 葡萄牙语 -> 中文
	opts.From = ydfanyi.PT
	res, _ = ydfanyi.Do("Olá, mundo.", opts)
	fmt.Println(res.String())

	// 意大利语 -> 中文
	opts.From = ydfanyi.IT
	res, _ = ydfanyi.Do("Il mondo intero", opts)
	fmt.Println(res.String())

	// 越南语 -> 中文
	opts.From = ydfanyi.VI
	res, _ = ydfanyi.Do("Xin chào thế giới", opts)
	fmt.Println(res.String())

	// 印尼语 -> 中文
	opts.From = ydfanyi.ID
	res, _ = ydfanyi.Do("Halo dunia", opts)
	fmt.Println(res.String())

	// 阿拉伯语 -> 中文
	opts.From = ydfanyi.AR
	res, _ = ydfanyi.Do("مرحبا العالم", opts)
	fmt.Println(res.String())

	// 荷兰语 -> 中文
	opts.From = ydfanyi.NL
	res, _ = ydfanyi.Do("Hallo wereld", opts)
	fmt.Println(res.String())
}
