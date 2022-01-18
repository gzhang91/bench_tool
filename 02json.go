package bench

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsonOperator struct {
	OperatorType OperatorType
	Resp         *http.Response
	Content      []byte
}

func NewJsonOperator() *JsonOperator {
	return &JsonOperator{
		OperatorType: OPERATOR_JSON,
	}
}

func (m *JsonOperator) DoGet(v interface{}) error {
	return nil
}

func (m *JsonOperator) DoPost(v interface{}) error {
	client := http.Client{}

	param, ok := v.(*RequestParam)
	if !ok {
		fmt.Printf("[DoPost] parse params faild, err: %v\n", ok)
		return fmt.Errorf("参数转换失败")
	}

	var req *http.Request
	var err error

	req, err = http.NewRequest("POST", param.Uri, bytes.NewReader([]byte(param.Body)))
	if err != nil {
		fmt.Printf("[DoPost] create request failed, err: %v", err)
		return err
	}

	req.Header = map[string][]string{}
	req.Header.Set("request", param.Request)

	headers, err := json.Marshal(param.Headers)
	if err != nil {
		fmt.Printf("[DoPost] parse headers faild, err: %v", err)
		return err
	}
	req.Header.Set("headers", string(headers))

	m.Resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("[DoPost] send post request failed, err: %s", err)
		return err
	}
	defer m.Resp.Body.Close()

	m.Content, err = ioutil.ReadAll(m.Resp.Body)
	if err != nil {
		fmt.Printf("[DoPost] parse response faild, err: %v", err)
		return err
	}

	return nil
}

func (m *JsonOperator) GetServiceType() OperatorType {
	return m.OperatorType
}
