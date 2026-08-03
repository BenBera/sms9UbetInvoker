package sms9UbetInvoker

import (
	"encoding/json"
	"fmt"
	"github.com/BenBera/sms9UbetInvoker/cryptor"
	"github.com/BenBera/sms9UbetInvoker/httpClient"
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
func (this Invoker) UserBet(_mobile string, _gameId string, _content SMSBetRequest) (error, *BetResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/bet")
	hc.SetMethod("POST")
	hc.SetHeaders("KEY-SHOPID", this.shopId)

	contentJSON, err := json.Marshal(_content)
	if err != nil {
		return fmt.Errorf("marshal SMS bet request: %w", err), nil
	}

	params := map[string]any{
		"mobile":    _mobile,
		"game_uuid": _gameId,
		"content":   string(contentJSON),
	}
	reqBody := this.Sign(params)
	hc.SetBody([]byte(cryptor.JsonEncode(map[string]any{
		"data": reqBody,
	})))

	err, resp := hc.Do()
	if err != nil {
		return err, nil
	}
	var betResp BetResp
	if err := json.Unmarshal([]byte(resp.Body), &betResp); err != nil {
		return fmt.Errorf("decode bet response: %w", err), nil
	}

	return err, &betResp
}
