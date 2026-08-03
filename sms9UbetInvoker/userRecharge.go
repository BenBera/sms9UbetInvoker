package sms9UbetInvoker

import (
	"encoding/json"
	"github.com/leo202465/sms9UbetInvoker/cryptor"
	"github.com/leo202465/sms9UbetInvoker/httpClient"
)

type UserRechargeResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type UserRechargeData struct {
	OrderId   string  `json:"order_id"`
	Price     float64 `json:"price"`
	RealPrice float64 `json:"real_price"`
	Tax       float64 `json:"tax"`
}

func (this UserRechargeResp) GetData() UserRechargeData {
	if this.Code > 0 {
		return UserRechargeData{}
	}
	dataJson := cryptor.JsonEncode(this.Data)
	var dataObj UserRechargeData
	err := json.Unmarshal([]byte(dataJson), &dataObj)
	if err != nil {
		return UserRechargeData{}
	}
	return dataObj
}

// 用户充值
func (this Invoker) UserRecharge(_mobile string, _price float64) (error, *UserRechargeResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/recharge")
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
	var respObj UserRechargeResp
	if err := resp.ToObj(&respObj); err != nil {
		return err, nil
	}
	return err, &respObj
}
