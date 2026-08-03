package sms9UbetInvoker

import (
	"encoding/json"
	"fmt"
	"github.com/leo202465/sms9UbetInvoker/cryptor"
	"github.com/leo202465/sms9UbetInvoker/httpClient"
)

type GetAccountBalanceResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type GetAccountBalanceData struct {
	Balance float64 `json:"balance"`
	Id      uint32  `json:"id"`
	Mobile  string  `json:"mobile"`
}

func (this GetAccountBalanceResp) GetData() GetAccountBalanceData {
	if this.Code > 0 {
		return GetAccountBalanceData{}
	}
	dataJson := cryptor.JsonEncode(this.Data)
	var dataObj GetAccountBalanceData
	err := json.Unmarshal([]byte(dataJson), &dataObj)
	if err != nil {
		return GetAccountBalanceData{}
	}
	return dataObj
}

// 获取账号余额
func (this Invoker) GetAccountBalance(_mobile string) (error, *GetAccountBalanceResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/getAccountBalance")
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
	fmt.Printf("resp: %v\n", string(resp.Body))
	var respObj GetAccountBalanceResp
	if err := resp.ToObj(&respObj); err != nil {
		return err, nil
	}
	return err, &respObj
}
