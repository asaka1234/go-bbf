package go_bbf

import (
	"encoding/json"
	"errors"
	"github.com/asaka1234/go-bbf/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

// 充值回调
func (cli *Client) DepositCallback(req BbfDepositBackReq, processor func(BbfDepositBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	// Verify signature
	flag, err := utils.Verify(params, cli.BackKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		return err
	}
	if !flag {
		reqJson, _ := json.Marshal(req)
		log.Printf("bbf#back#verify#fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
