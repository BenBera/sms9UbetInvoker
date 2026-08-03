package sms9UbetInvoker

import (
	"encoding/json"
	"github.com/leo202465/sms9UbetInvoker/cryptor"
	"github.com/leo202465/sms9UbetInvoker/httpClient"
)

type GetUserBetLogsResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type UserBetLogsData struct {
	List  []UserBetInfo `json:"list"`
	Total struct {
		TotalBetAmount float64 `json:"total_bet_amount"`
		TotalValidBet  float64 `json:"total_valid_bet"`
		TotalNetResult float64 `json:"total_net_result"`
		TotalWinAmount float64 `json:"total_win_amount"`
		TotalCount     uint32  `json:"total_count"`
		TotalTax       float64 `json:"total_tax"`
	} `json:"total"`
}

type UserBetInfo struct {
	Id             uint32  `json:"id"`
	User           string  `json:"user"`
	OrderNo        string  `json:"order_no"`
	BetAmount      float64 `json:"bet_amount"`
	ValidBet       float64 `json:"valid_bet"`
	WinAmount      float64 `json:"win_amount"`
	PlatId         uint32  `json:"plat_id"`
	PlatName       string  `json:"plat_name"`
	GameType       uint32  `json:"game_type"`
	CreateTime     string  `json:"create_time"`
	SettlementTime string  `json:"settlement_time"`
	Currency       string  `json:"currency"`
	Tax            float64 `json:"tax"`
	Status         string  `json:"status"`
	NetResult      float64 `json:"net_result"`
}

func (this GetUserBetLogsResp) GetData() UserBetLogsData {
	if this.Code > 0 {
		return UserBetLogsData{}
	}
	dataJson := cryptor.JsonEncode(this.Data)
	var dataObj UserBetLogsData
	err := json.Unmarshal([]byte(dataJson), &dataObj)
	if err != nil {
		return UserBetLogsData{}
	}
	return dataObj
}

// 用户充值
type GetUserBetLogsQuery struct {
	StartTime string // Start Time, format: "YYYY-MM-DD HH:ii:ss"
	EndTime   string // End Time, format: "YYYY-MM-DD HH:ii:ss"
	Page      uint32 // Page id, from 0 to 100, default 0
	Count     uint32 // Count of per page, default 20, can not be 0, and cant larger than 100
}

func (this Invoker) GetUserBetLogs(_mobile string, _queryOpt GetUserBetLogsQuery) (error, *GetUserBetLogsResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/getUserBetLogs")
	hc.SetMethod("POST")
	hc.SetHeaders("KEY-SHOPID", this.shopId)
	if _queryOpt.Count == 0 {
		_queryOpt.Count = 20
	}
	if _queryOpt.Page > 100 {
		_queryOpt.Page = 100
	}
	params := map[string]any{
		"mobile":     _mobile,
		"begin_time": _queryOpt.StartTime,
		"end_time":   _queryOpt.EndTime,
		"page":       _queryOpt.Page,
		"count":      _queryOpt.Count,
	}
	reqBody := this.Sign(params)
	hc.SetBody([]byte(cryptor.JsonEncode(map[string]any{
		"data": reqBody,
	})))

	err, resp := hc.Do()
	if err != nil {
		return err, nil
	}
	var respObj GetUserBetLogsResp
	if err := resp.ToObj(&respObj); err != nil {
		return err, nil
	}
	return err, &respObj
}
