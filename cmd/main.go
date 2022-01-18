package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	bench "waf_learn_bench"
)

func main() {
	flag.IntVar(&bench.GlobalConfig.TestType, "t", 1, "进行测试的类型，1为fact模式，2为json模式")
	flag.IntVar(&bench.GlobalConfig.SpecialIndex, "s", 0, "指定只测试某个index的数据（单测使用）")
	flag.StringVar(&bench.GlobalConfig.InputFilename, "i", "./", "输入测试的文件名路径")
	flag.StringVar(&bench.GlobalConfig.Uri, "u", "http://192.168.12.206:18081/json", "waf深度学习服务地址")
	flag.IntVar(&bench.GlobalConfig.NTimes, "n", 1, "执行压测总的次数")
	flag.IntVar(&bench.GlobalConfig.NRequests, "r", 8, "执行压测总的请求数")
	flag.IntVar(&bench.GlobalConfig.NConcurrency, "c", 1, "执行压测并行连接数")
	flag.IntVar(&bench.GlobalConfig.NTimeout, "T", 8, "连接深度学习服务超时时间")
	flag.Parse()

	if bench.GlobalConfig.Uri == "" {
		fmt.Printf("参数uri必须设置，现在为: %s\n", bench.GlobalConfig.Uri)
		return
	}

	if bench.GlobalConfig.NConcurrency > bench.GlobalConfig.NRequests {
		fmt.Printf("参数设置失败，requests不能小于concurrency\n")
		return
	}

	if bench.GlobalConfig.SpecialIndex < 0 {
		bench.GlobalConfig.SpecialIndex = 0
	}

	file_path := bench.GlobalConfig.InputFilename
	file_list, err := ioutil.ReadDir(file_path)
	if err != nil {
		fmt.Printf("扫描目录失败，错误为: %v\n", err)
		return
	}

	buffers := make([]string, 0)
	line_cnt := 0
	for idx := range file_list {
		if file_list[idx].Name() == "." || file_list[idx].Name() == ".." {
			continue
		}

		file, err := os.Open(file_path + "/" + file_list[idx].Name())
		if err != nil {
			fmt.Printf("打开文件[%s]失败，错误为: %v\n", file_list[idx].Name(), err)
			return
		}
		defer file.Close()

		reader := bufio.NewReader(file)

		var buf []byte
		for {
			line, prefix, err := reader.ReadLine()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("打开文件[%s]失败，错误为: %v\n", file_list[idx].Name(), err)
				return
			}

			buf = append(buf, line...)

			if prefix {
				continue
			}

			line_cnt++
			buffers = append(buffers, string(buf))
			buf = append(buf[:0], buf[len(buf):]...)
		}
	}

	bench.GlobalConfig.FileLines = line_cnt

	if bench.GlobalConfig.SpecialIndex > line_cnt {
		bench.GlobalConfig.SpecialIndex = 1
	}

	bench.GlobalConfig.ToString()

	MakeBenchTest(buffers)
}

func MakeBenchTest(buffers []string) {
	bench.NewBench(buffers).MakeBenchTest()
}
