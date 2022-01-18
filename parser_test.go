package bench

import (
	"fmt"
	"testing"
)

func TestParseFrom(t *testing.T) {
	req_str := `{"request":"GET /jxggzy/services/JyxxWebservice/getTradeList?response=application/json\u0026pageIndex=1\u0026pageSize=100\u0026\u0026dsname=ztb_data\u0026bname=\u0026qytype=3\u0026itemvalue=132 HTTP/1.1","headers":{"accept":"*/*","accept-encoding":"gzip, deflate","host":"www.fake.com","user-agent":"Python/3.9 aiohttp/3.7.4.post0"},"body":"-1","waf_info":[{"payload":"Python/3.9 aioht","zone":"HEADERS","type":"Crawler"}]}`
	req_obj := &RequestItem{}

	err := req_obj.ParseFrom(req_str)
	if err != nil {
		fmt.Printf("parse faild, err: %v", err)
	}

}
