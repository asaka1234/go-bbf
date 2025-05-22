package go_bbf

// ----------pre generate-------------------------
type BbfDepositReq struct {
	UniqueCode string `json:"uniqueCode" mapstructure:"uniqueCode"`
	Money      string `json:"money" mapstructure:"money"`
	PayType    int    `json:"payType" mapstructure:"payType"`
	OrderId    string `json:"orderId" mapstructure:"orderId"`
	PayerName  string `json:"payerName" mapstructure:"payerName"`
	//这个不需要业务侧使用,而是sdk帮计算和补充
	//Signature string `json:"signature"`
	//UID       int    `json:"uid"`
}

type BbfDepositRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

//==============================================

type BbfDepositBackReq struct {
	ApiOrderNo  string `json:"apiOrderNo" mapstructure:"money"`
	Money       string `json:"money" mapstructure:"money"`
	TradeStatus int    `json:"tradeStatus" mapstructure:"tradeStatus"`
	TradeId     string `json:"tradeId" mapstructure:"tradeId"`
	UniqueCode  string `json:"uniqueCode" mapstructure:"uniqueCode"`
	Signature   string `json:"signature" mapstructure:"signature"`
}

type BbfDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}
