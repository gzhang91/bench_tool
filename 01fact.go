package bench

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FactOperator struct {
	OperatorType OperatorType
	Resp         *http.Response
	Content      []byte
}

func NewFactOperator() *FactOperator {
	return &FactOperator{
		OperatorType: OPERATOR_FACT,
	}
}

func (m *FactOperator) DoGet(v interface{}) error {
	return nil
}

func (m *FactOperator) DoPost(v interface{}) error {
	param, ok := v.(*RequestParam)
	if !ok {
		fmt.Printf("[DoPost] parse params faild, err: %v\n", ok)
		return fmt.Errorf("参数转换失败")
	}

	resp, err := http.PostForm(param.Uri,
		url.Values{"inputStr": {param.Request}})

	if err != nil {
		fmt.Printf("[DoPost] create request failed, err: %v", err)
		return err
	}

	defer resp.Body.Close()

	m.Resp = resp
	m.Content, err = ioutil.ReadAll(m.Resp.Body)
	if err != nil {
		fmt.Printf("[DoPost] parse response faild, err: %v", err)
		return err
	}

	return nil
}

func (m *FactOperator) GetServiceType() OperatorType {
	return m.OperatorType
}
