package sms9UbetInvoker

import (
	"encoding/json"
	"fmt"
	"github.com/leo202465/sms9UbetInvoker/cryptor"
	"github.com/leo202465/sms9UbetInvoker/httpClient"
)

type GetGameListResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type GetGameListData struct {
	List []struct {
		GameUUID string `json:"game_uuid"`
		GameName string `json:"game_name"`
	} `json:"list"`
	Total uint32 `json:"total"`
}

func (this GetGameListResp) GetData() GetGameListData {
	if this.Code > 0 {
		return GetGameListData{}
	}
	dataJson := cryptor.JsonEncode(this.Data)
	var dataObj GetGameListData
	err := json.Unmarshal([]byte(dataJson), &dataObj)
	if err != nil {
		return GetGameListData{}
	}
	return dataObj
}

// 获取游戏列表
func (this Invoker) GetGameList(_page int) (error, *GetGameListResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/getGameList")
	hc.SetMethod("POST")
	hc.SetHeaders("KEY-SHOPID", this.shopId)

	params := map[string]any{
		"page": _page,
	}
	reqBody := this.Sign(params)
	hc.SetBody([]byte(cryptor.JsonEncode(map[string]any{
		"data": reqBody,
	})))

	err, resp := hc.Do()
	if err != nil {
		return err, nil
	}
	fmt.Printf("RESP: %v\n", string(resp.Body))
	var respObj GetGameListResp
	if err := resp.ToObj(&respObj); err != nil {
		return err, nil
	}
	return err, &respObj
}
