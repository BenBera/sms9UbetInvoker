package sms9UbetInvoker

import (
	"encoding/json"
	"github.com/leo202465/sms9UbetInvoker/cryptor"
	"github.com/leo202465/sms9UbetInvoker/httpClient"
)

type RegAccountResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type RegAccountData struct {
	Balance float64 `json:"balance"`
	Id      uint32  `json:"id"`
}

func (this RegAccountResp) GetData() RegAccountData {
	if this.Code > 0 {
		return RegAccountData{}
	}
	dataJson := cryptor.JsonEncode(this.Data)
	var dataObj RegAccountData
	err := json.Unmarshal([]byte(dataJson), &dataObj)
	if err != nil {
		return RegAccountData{}
	}
	return dataObj
}

// 注册账号
func (this Invoker) RegAccount(_mobile string) (error, *RegAccountResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/regAccount")
	hc.SetMethod("POST")
	hc.SetHeaders("KEY-SHOPID", this.shopId)

	params := map[string]any{
		"mobile": _mobile,
	}
	reqBody := this.Sign(params)
	hc.SetBody([]byte(cryptor.JsonEncode(map[string]any{
		"data": reqBody,
	})))

	err, resp := hc.Do()
	if err != nil {
		return err, nil
	}
	var respObj RegAccountResp
	if err := resp.ToObj(&respObj); err != nil {
		return err, nil
	}
	return err, &respObj
}
