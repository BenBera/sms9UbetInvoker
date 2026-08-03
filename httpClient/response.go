package httpClient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 回应结构体
type Response struct {
	// 回应头
	Header http.Header
	// 回应体
	Body []byte
	// 状态码
	StatusCode uint32
	// 回应体长度
	ContentLength uint32
}

// 生成新的回应体
func NewResponse(_resp *http.Response) *Response {
	var resp = Response{
		Header:        _resp.Header,
		Body:          []byte{},
		StatusCode:    uint32(_resp.StatusCode),
		ContentLength: uint32(_resp.ContentLength),
	}
	var buffer, err = ioutil.ReadAll(_resp.Body)
	if err != nil {
		return nil
	}
	resp.Body = buffer
	return &resp
}

func (this *Response) ToObj(_obj any) error {
	return json.Unmarshal(this.Body, _obj)
}
