package httpClient

import (
	"net"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	Url     string
	Method  string
	Body    []byte
	headers map[string]string
	timeout time.Duration
}

// 创建HTTP客户端
func NewClient(_url string) HttpClient {
	return HttpClient{
		Url:    _url,
		Method: "POST",
		headers: map[string]string{
			"Content-Type": "text/json",
		},
		Body:    make([]byte, 0),
		timeout: 30 * time.Second,
	}
}

func (this *HttpClient) SetMethod(_method string) {
	this.Method = _method
}

func (this *HttpClient) SetBody(_body []byte) {
	this.Body = _body
}

func (this *HttpClient) SetTimeout(_time time.Duration) {
	this.timeout = _time
}

func (this *HttpClient) SetHeaders(_key string, _val string) {
	this.headers[_key] = _val
}

func (this *HttpClient) Do() (error, *Response) {
	body := strings.NewReader(string(this.Body))
	req, err := http.NewRequest(this.Method, this.Url, body)
	if err != nil {
		return err, nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, this.timeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(this.timeout))
				return conn, nil
			},
			ResponseHeaderTimeout: this.timeout,
		},
	}

	// 添加头
	for HKey, HVal := range this.headers {
		req.Header.Set(HKey, HVal)
	}

	res, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	resp := NewResponse(res)
	res.Body.Close()
	return nil, resp
}
