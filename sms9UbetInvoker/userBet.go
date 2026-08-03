package sms9UbetInvoker

import (
	"encoding/json"
	"github.com/leo202465/sms9UbetInvoker/cryptor"
	"github.com/leo202465/sms9UbetInvoker/httpClient"
)

type BetResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type BetData struct {
	OrderId string  `json:"order_id"` // 订单号
	Balance float64 `json:"balance"`  // 下注后余额
	Amount  float64 `json:"amount"`   // 下注金额
}

func (this BetResp) GetData() BetData {
	if this.Code > 0 {
		return BetData{}
	}
	dataJson := cryptor.JsonEncode(this.Data)
	var dataObj BetData
	err := json.Unmarshal([]byte(dataJson), &dataObj)
	if err != nil {
		return BetData{}
	}
	return dataObj
}

// 获取游戏列表
func (this Invoker) UserBet(_mobile string, _gameId string, _content string) (error, *BetResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/bet")
	hc.SetMethod("POST")
	hc.SetHeaders("KEY-SHOPID", this.shopId)

	params := map[string]any{
		"mobile":    _mobile,
		"game_uuid": _gameId,
		"content":   _content,
	}
	reqBody := this.Sign(params)
	hc.SetBody([]byte(cryptor.JsonEncode(map[string]any{
		"data": reqBody,
	})))

	err, resp := hc.Do()
	if err != nil {
		return err, nil
	}
	var respObj BetResp
	if err := resp.ToObj(&respObj); err != nil {
		return err, nil
	}
	return err, &respObj
}
