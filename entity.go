package go_bbf

// ----------pre generate-------------------------
type BbfDepositReq struct {
	UID        int    `json:"uid"`
	UniqueCode string `json:"uniqueCode"`
	Money      string `json:"money"`
	PayType    int    `json:"payType"`
	OrderId    string `json:"orderId"`
	Signature  string `json:"signature"`
	PayerName  string `json:"payerName"`
}

type BbfDepositRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

type BbfDepositBackReq struct {
	ApiOrderNo  string `json:"apiOrderNo"`
	Money       string `json:"money"`
	TradeStatus int    `json:"tradeStatus"`
	TradeId     string `json:"tradeId"`
	UniqueCode  string `json:"uniqueCode"`
	Signature   string `json:"signature"`
}

type BbfDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}
