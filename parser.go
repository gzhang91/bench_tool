package bench

import (
	"encoding/json"
	"log"
)

/*
sample:
	{
	"request": "GET /taxapp/resource/js/utilsNoneJqueryMobile.js?v=20201116 HTTP/1.1",
	"headers": {
		"accept": "*\/*",
		"accept-encoding": "gzip, deflate",
		"accept-language": "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7",
		"connection": "keep-alive",
		"cookie": "SESSION=9cad3aaf-8603-47c6-8d4a-ffab6513097f",
		"host": "www.fake.com",
		"referer": "http://www.fake.com/taxapp/pages/person/main.jsp",
		"user-agent": "Mozilla/5.0 (Linux; Android 11; V1955A Build/RP1A.200720.012; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.106 Mobile Safari/537.36",
		"x-requested-with": "com.powersi.zhsw"
		},
		"body": "-1"
	}
*/

type RequestItem struct {
	RequestLine string            `json:"request"`
	Headers     map[string]string `json:"headers"`
	Body        string            `json:"body"`
	// WafInfo     []map[string]string `json:"waf_info"`
}

func NewRequestItem() *RequestItem {
	return &RequestItem{}
}

func (m *RequestItem) ParseFrom(request_str string) error {
	err := json.Unmarshal([]byte(request_str), m)
	if err != nil {
		log.Fatalf("%s 解析request字符串失败，错误为: %v", LOG_TITLE, err)
		return err
	}

	return err
}

func (m *RequestItem) TransforTo() (string, error) {
	res, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("%s 序列化request失败，错误为: %v", LOG_TITLE, err)
		return "", nil
	}

	return string(res), nil
}
