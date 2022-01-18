package bench

import "fmt"

var (
	GlobalConfig Config
)

type Config struct {
	TestType      int    // 测试类型，是fact，还是json
	SpecialIndex  int    // 指定只测试某个index的数据（单测使用）
	InputFilename string // 输入文件路径
	Uri           string // 要访问的Uri
	NTimes        int    // 执行次数
	NRequests     int    // 总的发送的请求数
	NConcurrency  int    // 并行连接数
	NTimeout      int    // 访问超时时间
	FileLines     int    // 文件行数
}

func (m *Config) ToString() {
	fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
	fmt.Printf("++参数配置如下: \n")
	fmt.Printf("+ 测试类型为: %v\n", m.TestType)
	fmt.Printf("+ 输入文件名: %v\n", m.InputFilename)
	fmt.Printf("+ 访问的URI: %v\n", m.Uri)
	fmt.Printf("+ 执行次数: %v\n", m.NTimes)
	fmt.Printf("+ 发送的请求数: %v\n", m.NRequests)
	fmt.Printf("+ 并发数: %v\n", m.NConcurrency)
	fmt.Printf("+ 超时: %v\n", m.NTimeout)
	fmt.Printf("+ 文件行数: %v\n", m.FileLines)
	fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++++++++++++\n\n")
}
