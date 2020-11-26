package ydfanyi

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const Version = "0.0.1"

const (
	api       = "http://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule"
	referer   = "http://fanyi.youdao.com/"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0"
	cookie    = "OUTFOX_SEARCH_USER_ID=-722430601@10.169.0.102"
)

const (
	AUTO = "AUTO"   // 自适应
	ZH   = "zh-CHS" // 汉语
	EN   = "en"     // 英语
	JA   = "ja"     // 日语
	KO   = "ko"     // 汉语
	FR   = "fr"     // 法语
	DE   = "de"     // 德语
	RU   = "ru"     // 俄语
	ES   = "es"     // 西班牙语
	PT   = "pt"     // 葡萄牙语
	IT   = "it"     // 意大利语
	VI   = "vi"     // 越南语
	ID   = "id"     // 印尼语
	AR   = "ar"     // 阿拉伯语
	NL   = "nl"     // 荷兰语
)

type Result struct {
	Code int `json:"errorCode"`
	Data [][]struct {
		Tgt string `json:"tgt"`
		Src string `json:"src"`
	} `json:"translateResult"`
	Type string `json:"type"`
	SmartResult struct{
		Entries []string `json:"entries"`
	} `json:"smartResult"`
}

func (r Result) String() string {
	var out string
	if len(r.Data) > 0 && len(r.Data[0]) > 0 {
		out = r.Data[0][0].Tgt
	}

	return out
}

type Options struct {
	Proxy string
	From  string
	To    string
}

func NewOptions(from, to, proxy string) *Options {
	if from == "" {
		from = AUTO
	}

	if to == "" {
		to = AUTO
	}

	return &Options{
		Proxy: proxy,
		From:  from,
		To:    to,
	}
}

func Do(str string, options *Options) (result Result, err error) {
	if options == nil {
		options = NewOptions(AUTO, AUTO, "")
	}

	if options.From != AUTO && options.To != AUTO && options.From != ZH && options.To != ZH {
		err = fmt.Errorf("%s -> %s translation is not supported", options.From, options.To)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(*http.Request) (*url.URL, error) {
				if options.Proxy == "" {
					return nil, nil
				}
				return url.ParseRequestURI(options.Proxy)
			},
			DisableCompression:  true,
			TLSHandshakeTimeout: 10 * time.Second,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
	}

	rand.Seed(time.Now().Unix())
	t := strconv.Itoa(int(time.Now().Unix()))
	t += fmt.Sprintf("%d%d%d%d", rand.Intn(9), +rand.Intn(9), +rand.Intn(9), +rand.Intn(9))

	body := strings.NewReader(url.Values{
		"i":           []string{str},
		"from":        []string{options.From},
		"to":          []string{options.To},
		"smartresult": []string{"dict"},
		"client":      []string{"fanyideskweb"},
		"salt":        []string{t},
		"sign":        []string{sign(fmt.Sprintf("fanyideskweb%s%s]BjuETDhU)zqSxf-=B#7m", str, t))},
		"doctype":     []string{"json"},
		"version":     []string{"2.1"},
		"keyfrom":     []string{"fanyi.web"},
	}.Encode())

	var req *http.Request
	if req, err = http.NewRequest("POST", api, body); err != nil {
		return
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", referer)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	if res.StatusCode >= 300 {
		err = fmt.Errorf(string(bytes))
		return
	}

	if err = json.Unmarshal(bytes, &result); err != nil {
		return
	}

	if result.Code != 0 {
		err = fmt.Errorf(string(bytes))
		return
	}

	return
}

func sign(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
