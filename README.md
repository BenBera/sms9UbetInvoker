```go
package main
// define your env variables
const shopId = "{YOUR_SHOP_ID}"
const AESKey = "{YOUR_AES_KEY}"
const MD5Key = "{YOUR_MD5_KEY}"
const ApiUrl = "{YOUR_API_URL}"

// new invoker
invoker := NewInvoker(shopId, AESKey, MD5Key, ApiUrl)

// do anything
// example, register a mobile
invoker.RegAccount("012345678")
```

### APIList
```go
    // join mobile to system
    RegAccount(_mobile string)
```
> Possible Resp:
- Code: 1008, Msg: mobile_invalid
  means mobile format is invalid, please use valid mobile number
- Code: 1010, Msg: mobile_is_exist
  means mobile is exist, please use another mobile number
	
```go
    // payment
    UserRecharge(_mobile string, _price float64)
```
> Possible Resp:
- Code: 1008, Msg: mobile_invalid
  means mobile format is invalid, please use valid mobile number
- Code: 10000, Msg: price_invalid
  means the price can not be 0 or null
- Code: 40001, Msg: account_not_found
  means the mobile is not found, please register first
- Code: 10006, Msg: account_permission_error
  means current account has no permission to do this
- Code: 10005, Msg: shop_not_found
  means current system has no valid payment shop, please contact manager to fix it.
- Code: 10009, Msg: must_be_large_than_min_price
  means the price must large than min price, the min price is in data param of response 
- Code: 10009, Msg: must_be_small_than_max_price
  means the price must less than max price, the max price is in data param of response


```go
    // withdraw
    UserWithdraw(_mobile string, _price float64)
```
> Possible Resp:
- Code: 1008, Msg: mobile_invalid
  means mobile format is invalid, please use valid mobile number
- Code: 10000, Msg: price_invalid
  means the price can not be 0 or null
- Code: 40001, Msg: account_not_found
  means the mobile is not found, please register first
- Code: 20000, Msg: account_frozen_contact_customer
  means current account is banned, please contact customer
- Code: 10005, Msg: shop_not_found
  means current system has no valid payment shop, please contact manager to fix it.
- Code: 10064, Msg: cant_withdraw
  means current account cant do withdraw this moment
- Code: 10009, Msg: resp_price_out_range
  means the price must be within the specified range, which is provided in the data field of the response.
- Code: 10066, Msg: resp_price_examine_req
  the current audit requirements have not been met, so withdrawal cannot be processed. Please complete the required audit turnover first.
- Code: 10068, Msg: balance_not_enough
  means current account has no enough balance to do this.
- Code: 10050, Msg: resp_withdraw_count_limit
  You have reached the withdrawal limit for today. Please try again tomorrow.
```go
    // get game list
    // page from 0 to 1000, 20 records / page
    GetGameList(_page uint32)
```
> Possible Resp:
- None


```go
    // bet to system
    // mobile: who bet
    // gameId: which game id from get game list api
    // content: the sms content
    UserBet(_mobile string, _gameId string, _content string)
```
> Possible Resp:
- Testing , Im not sure

```go
    // get user recharge logs
    /**
    type GetUserRechargeLogsQuery struct {
        StartTime string // Start Time, format: "YYYY-MM-DD HH:ii:ss"
        EndTime   string // End Time, format: "YYYY-MM-DD HH:ii:ss"
        Page      uint32 // Page id, from 0 to 100, default 0
        Count     uint32 // Count of per page, default 20, can not be 0, and cant larger than 100
    }
     */
    GetUserRechargeLogs(_mobile string, _queryOpt GetUserRechargeLogsQuery)
```
> Possible Resp:
- None

General error message:
- Code: 40000, Msg: system_error
  means something goes wrong, contact technology support.


```go
    // get user bet logs
    /**
    type GetUserBetLogsQuery struct {
        StartTime string // Start Time, format: "YYYY-MM-DD HH:ii:ss"
        EndTime   string // End Time, format: "YYYY-MM-DD HH:ii:ss"
        Page      uint32 // Page id, from 0 to 100, default 0
        Count     uint32 // Count of per page, default 20, can not be 0, and cant larger than 100
    }
     */
    GetUserBetLogs(_mobile string, _queryOpt GetUserRechargeLogsQuery)
```
> Possible Resp:
- None

General error message:
- Code: 40000, Msg: system_error
  means something goes wrong, contact technology support.


All Done, Good Job!!
