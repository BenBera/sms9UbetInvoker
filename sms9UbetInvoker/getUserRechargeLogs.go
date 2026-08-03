package sms9UbetInvoker

import (
	"encoding/json"
	"github.com/leo202465/sms9UbetInvoker/cryptor"
	"github.com/leo202465/sms9UbetInvoker/httpClient"
)

type GetUserRechargeLogsResp struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type UserRechargeLogsData struct {
	List  []UserRechargeInfo `json:"list"`
	Total struct {
		Count        uint32  `json:"count"`
		Recharge     float64 `json:"recharge"`
		RealRecharge float64 `json:"real_recharge"`
		Tax          float64 `json:"tax"`
	} `json:"total"`
}

type UserRechargeInfo struct {
	Id            uint32  `json:"id"`
	TradNo        string  `json:"trad_no"`
	Price         float64 `json:"price"`
	RealPrice     float64 `json:"real_price"`
	Tax           float64 `json:"tax"`
	CreateTime    string  `json:"create_time"`
	CompletedTime string  `json:"completed_time"`
	Status        uint32  `json:"status"`
}

func (this GetUserRechargeLogsResp) GetData() UserRechargeLogsData {
	if this.Code > 0 {
		return UserRechargeLogsData{}
	}
	dataJson := cryptor.JsonEncode(this.Data)
	var dataObj UserRechargeLogsData
	err := json.Unmarshal([]byte(dataJson), &dataObj)
	if err != nil {
		return UserRechargeLogsData{}
	}
	return dataObj
}

// 用户充值
type GetUserRechargeLogsQuery struct {
	StartTime string // Start Time, format: "YYYY-MM-DD HH:ii:ss"
	EndTime   string // End Time, format: "YYYY-MM-DD HH:ii:ss"
	Page      uint32 // Page id, from 0 to 100, default 0
	Count     uint32 // Count of per page, default 20, can not be 0, and cant larger than 100
}

func (this Invoker) GetUserRechargeLogs(_mobile string, _queryOpt GetUserRechargeLogsQuery) (error, *GetUserRechargeLogsResp) {
	hc := httpClient.NewClient(this.baseUrl + "/do/getUserRechargeLogs")
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
	var respObj GetUserRechargeLogsResp
	if err := resp.ToObj(&respObj); err != nil {
		return err, nil
	}
	return err, &respObj
}
