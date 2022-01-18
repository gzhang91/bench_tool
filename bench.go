package bench

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg       = &sync.WaitGroup{}
	totalRes []TotalResult
)

type Bench struct {
	buffers []string
}

func NewBench(buf []string) *Bench {
	return &Bench{
		buffers: buf,
	}
}

func (m *Bench) MakeBenchTest() {
	try := 0
	per_reqs := GlobalConfig.NRequests / GlobalConfig.NConcurrency

	for try < GlobalConfig.NTimes {
		fmt.Printf("第[%d]次压测 =============================================================================== \n", try+1)

		totalRes = []TotalResult{}
		totalRes = make([]TotalResult, GlobalConfig.NConcurrency)
		concurr := 0
		for concurr < GlobalConfig.NConcurrency {
			wg.Add(1)
			// fmt.Printf("index为: %v\n", concurr)

			if GlobalConfig.TestType == 1 {
				go m.makeRequestFact(concurr, per_reqs)
			} else {
				go m.makeRequestJson(concurr, per_reqs)
			}

			concurr++
		}

		wg.Wait()

		tos := TotalResult{
			StartTime:       0,
			EndTime:         0,
			SuccessRequests: 0,
			TotalRequests:   0,
			FailedRate:      0,
		}
		for _, val := range totalRes {
			// fmt.Printf("startTime: %v, endTime: %v\n", val.StartTime, val.EndTime)

			if tos.StartTime == 0 || tos.StartTime-val.StartTime > 0 {
				tos.StartTime = val.StartTime
			}

			if tos.EndTime == 0 || tos.EndTime-val.EndTime < 0 {
				tos.EndTime = val.EndTime
			}

			tos.TotalRequests += val.TotalRequests
			tos.SuccessRequests += val.SuccessRequests
			tos.AverageTime += val.AverageTime
			tos.FailedRate += val.FailedRate
			tos.Concurrency = val.Concurrency
			tos.AbormalCnt += val.AbormalCnt
			tos.NormalCnt += val.NormalCnt
		}

		tos.AverageTime = int64(float64(tos.AverageTime) / float64(tos.SuccessRequests)) // ms
		tos.ToString()

		fmt.Printf("第[%d]次压测完毕 ===========================================================================\n\n", try+1)
		try++
	}

}

func (m *Bench) makeRequestFact(index int, max_reqs int) {
	defer wg.Done()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	req := 0
	rindex := r.Int63n(int64(len(m.buffers) - 1))
	//fmt.Printf("index: %v\n", rindex)
	for req < max_reqs {
		req++

		if GlobalConfig.SpecialIndex > 0 {
			rindex = int64(GlobalConfig.SpecialIndex - 1)
		} else {
			rindex++
			rindex = rindex % int64(len(m.buffers))
		}

		rp := RequestParam{
			Uri:     GlobalConfig.Uri,
			Request: m.buffers[rindex],
		}

		start := time.Now()
		if totalRes[index].StartTime == 0 {
			totalRes[index].StartTime = start.UnixMilli()
		}

		jp := NewFactOperator()
		err := jp.DoPost(&rp)
		if err != nil {
			fmt.Printf("[makeRequest] 发送post请求失败，错误为: %v\n", err)
			continue
		}

		if jp.Resp.StatusCode == 200 {
			result := Result{}
			err := result.ParseFrom(string(jp.Content))
			if err != nil {
				fmt.Printf("[makeRequest] 解析post响应失败，错误为: %v\n", err)
				continue
			}

			totalRes[index].SuccessRequests++
			if result.Data.Result == 0 {
				totalRes[index].NormalCnt++
			}

			if result.Data.Result == 1 {
				totalRes[index].AbormalCnt++
			}

			end := time.Now()
			if end.UnixMilli() > start.UnixMilli() { // 排除异常数据
				totalRes[index].AverageTime += (end.UnixMilli() - start.UnixMilli())
			}
			// fmt.Printf("begin: %v, end: %v, delta: %v\n", end.UnixMilli(), start.UnixMilli(), end.UnixMilli()-start.UnixMilli())
		} else {
			fmt.Printf("[makeRequest] 发送post请求失败，错误码为: %v\n", jp.Resp.StatusCode)
			continue
		}

		totalRes[index].TotalRequests++
	}

	totalRes[index].EndTime = time.Now().UnixMilli()
	totalRes[index].Concurrency = int32(GlobalConfig.NConcurrency)

	// fmt.Printf("succ: %v, total: %v\n", totalRes[index].SuccessRequests, totalRes[index].TotalRequests)
}

func (m *Bench) makeRequestJson(index int, max_reqs int) {
	defer wg.Done()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	req := 0

	rindex := r.Int63n(int64(len(m.buffers) - 1))

	// fmt.Printf("index: %v\n", rindex)
	for req < max_reqs {
		req++
		if GlobalConfig.SpecialIndex > 0 {
			rindex = int64(GlobalConfig.SpecialIndex - 1)
		} else {
			rindex++
			rindex = rindex % int64(len(m.buffers))
		}

		ri := NewRequestItem()
		err := ri.ParseFrom(m.buffers[rindex])
		if err != nil {
			fmt.Printf("[makeRequest] 解析[%s]失败，错误为: %v\n", m.buffers[rindex], err)
			continue
		}

		rp := RequestParam{
			Uri:     GlobalConfig.Uri,
			Request: ri.RequestLine,
			Headers: ri.Headers,
			Body:    ri.Body,
		}

		start := time.Now()
		if totalRes[index].StartTime == 0 {
			totalRes[index].StartTime = start.UnixMilli()
		}

		jp := NewJsonOperator()
		err = jp.DoPost(&rp)
		if err != nil {
			fmt.Printf("[makeRequest] 发送post请求失败，错误为: %v\n", err)
			continue
		}

		if jp.Resp.StatusCode == 200 {
			result := Result{}
			err := result.ParseFrom(string(jp.Content))
			if err != nil {
				fmt.Printf("[makeRequest] 解析post响应失败，错误为: %v\n", err)
				continue
			}

			totalRes[index].SuccessRequests++
			if result.Data.Result == 0 {
				totalRes[index].NormalCnt++
			}

			if result.Data.Result == 1 {
				totalRes[index].AbormalCnt++
			}

			end := time.Now()
			if end.UnixMilli() > start.UnixMilli() { // 排除异常数据
				totalRes[index].AverageTime += (end.UnixMilli() - start.UnixMilli())
			}
		} else {
			fmt.Printf("[makeRequest] 发送post请求失败，错误码为: %v\n", jp.Resp.StatusCode)
			continue
		}

		totalRes[index].TotalRequests++
	}
	totalRes[index].EndTime = time.Now().UnixMilli()
	totalRes[index].Concurrency = int32(GlobalConfig.NConcurrency)
}
