package sms9UbetInvoker

type Invoker struct {
	aesKey  string
	md5Key  string
	baseUrl string
	shopId  string
}

type SMSBetRequest struct {
	Phone      string            `json:"phone"`
	Stake      float64           `json:"stake"`
	BetType    int               `json:"bet_type"`
	Selections []SMSBetSelection `json:"selections"`
}

type SMSBetSelection struct {
	GameID     int    `json:"game_id"`
	MarketType string `json:"market_type"`
	Pick       string `json:"pick"`
}
