package bench

import (
	"encoding/json"
	"fmt"
	"time"
)

type TotalResult struct {
	StartTime       int64   // 第一个请求到达的时间
	EndTime         int64   // 最后一个请求到达的时间
	AverageTime     int64   // 平均时间
	TotalRequests   int64   // 总的请求数TotalRequest
	SuccessRequests int64   // 成功的请求数
	FailedRate      float64 // 失败率
	Concurrency     int32   // 并发度

	AbormalCnt int64 // 异常请求个数
	NormalCnt  int64 // 正常请求个数
}

func (m *TotalResult) ToString() {
	fmt.Printf("## 结果如下：\n")
	fmt.Printf("# 第一个请求到达时间：[%v](%v)\n", time.UnixMilli(m.StartTime).Local(), m.StartTime)
	fmt.Printf("# 最后一个请求到达时间：[%v](%v)\n", time.UnixMilli(m.EndTime).Local(), m.EndTime)
	fmt.Printf("# 请求的时延: %v ms\n", m.EndTime-m.StartTime)
	fmt.Printf("# 请求的平均时间: %v ms\n", m.AverageTime)
	fmt.Printf("# 总请求数: %v\n", m.TotalRequests)
	fmt.Printf("# 成功请求数: %v\n", m.SuccessRequests)
	fmt.Printf("# 请求失败率: %.3f\n", float32(m.TotalRequests-m.SuccessRequests)/float32(m.TotalRequests))
	fmt.Printf("# 并发度为: %v\n", m.Concurrency)
	fmt.Printf("# waf异常请求个数: %v\n", m.AbormalCnt)
	fmt.Printf("# waf正常请求个数: %v\n", m.NormalCnt)
}

type DataRes struct {
	Access   string `json:"access"`
	Unaccess string `json:"unaccess"`
	Result   int    `json:"result"`
}

type Result struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data DataRes `json:"data"`
}

func (m *Result) ParseFrom(param_str string) error {
	err := json.Unmarshal([]byte(param_str), m)
	if err != nil {
		fmt.Printf("%s 解析result字符串[%s]失败，错误为: %v\n", LOG_TITLE, param_str, err)
		return err
	}

	return err
}

func (m *Result) TransforTo() (string, error) {
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("%s 序列化result失败，错误为: %v\n", LOG_TITLE, err)
		return "", nil
	}

	return string(res), nil
}
