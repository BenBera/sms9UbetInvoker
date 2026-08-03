package sms9UbetInvoker

import (
	"fmt"
	"log"
	"testing"
)

/*
*
shop_id: 39845
aes_key: e2XEhtXtpgf2IFI4UCL3wUeMmZ7WKPX7
md5_key: gOvZM1ezAPYy8t0CcYnbR9jAvszfT6oB
base_url: https://sms-api-prod.9ubet.com
*/

//const shopId = "39845"
//const ApiUrl = "https://sms-api-prod.9ubet.com"
//const AESKey = "e2XEhtXtpgf2IFI4UCL3wUeMmZ7WKPX7"
//const MD5Key = "gOvZM1ezAPYy8t0CcYnbR9jAvszfT6oB"

const shopId = "49361"
const ApiUrl = "https://sms-api-staging.9ubet.com"
const AESKey = "VecsTwFCcg034ZsK8Oo7t0nWouxHJPgh"
const MD5Key = "9ahzccEk9EvKTPD67PC2Vp95Yz8sVMr5"

func TestInvoker_RegAccount(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.RegAccount("012345671")
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v", respObj.GetData())
}

func TestInvoker_GetAccountBalance(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.GetAccountBalance("012345671")
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}

func TestInvoker_UserRecharge(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.UserRecharge("012345671", 100)
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}

func TestInvoker_UserWithdraw(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.UserWithdraw("012345671", 100)
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}

func TestInvoker_GetGameList(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.GetGameList(0)
	if err != nil {
		log.Fatalf("Reg Account Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}

func TestInvoker_GetUserRechargeLogs(t *testing.T) {
	invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

	err, respObj := invoker.GetUserRechargeLogs("254708491516", GetUserRechargeLogsQuery{
		StartTime: "",
		EndTime:   "",
	})
	if err != nil {
		log.Fatalf("Get User Recharge Logs Error: %v", err)
		return
	}
	if respObj.Code > 0 {
		log.Fatalf("RespCode: %v - %v", respObj.Code, respObj.Msg)
		return
	}

	fmt.Printf("Resp: %+v\n", respObj.GetData())
}
