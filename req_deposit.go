package go_bbf

import (
	"crypto/tls"
	"encoding/json"
	"github.com/asaka1234/go-bbf/utils"
	"github.com/mitchellh/mapstructure"
)

// 下单
func (cli *Client) Deposit(req BbfDepositReq) (*BbfDepositRsp, error) {

	rawURL := cli.DepositURL

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["uid"] = cli.MerchantID

	// Generate signature
	signStr, _ := utils.Sign(params, cli.AccessKey)
	params["signature"] = signStr

	// Convert to JSON
	jsonStr, _ := json.Marshal(params)
	cli.logger.Infof("BbfCurService#deposit#json: %s", jsonStr)

	var result BbfDepositRsp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
