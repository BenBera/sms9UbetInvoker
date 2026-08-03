package sms9UbetInvoker

import (
	"encoding/json"
	"github.com/leo202465/sms9UbetInvoker/cryptor"
	"github.com/leo202465/sms9UbetInvoker/httpClient"
)

type UserWithdrawResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type UserWithdrawData struct {
	OrderId string  `json:"order_id"`
	Balance float64 `json:"balance"`
}

func (this UserWithdrawResp) GetData() UserWithdrawData {
	if this.Code > 0 {
		return UserWithdrawData{}
	}
	dataJson := cryptor.JsonEncode(this.Data)
	var dataObj UserWithdrawData
	err := json.Unmarshal([]byte(dataJson), &dataObj)
	if err != nil {
		return UserWithdrawData{}
	}
	return dataObj
}

// 用户提现
func (this Invoker) UserWithdraw(_mobile string, _price float64) (error, *UserWithdrawResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/withdraw")
	hc.SetMethod("POST")
	hc.SetHeaders("KEY-SHOPID", this.shopId)

	params := map[string]any{
		"mobile": _mobile,
		"price":  _price,
	}
	reqBody := this.Sign(params)
	hc.SetBody([]byte(cryptor.JsonEncode(map[string]any{
		"data": reqBody,
	})))

	err, resp := hc.Do()
	if err != nil {
		return err, nil
	}
	var respObj UserWithdrawResp
	if err := resp.ToObj(&respObj); err != nil {
		return err, nil
	}
	return err, &respObj
}
