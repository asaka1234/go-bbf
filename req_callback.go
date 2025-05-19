package go_bbf

import (
	"encoding/json"
	"errors"
	"github.com/asaka1234/go-bbf/utils"
	"log"
)

// 充值回调
func (cli *Client) DepositCallback(req BbfDepositBackReq, processor func(BbfDepositBackReq) error) error {
	//验证签名
	key := cli.BackKey

	// Convert struct to map[string]interface{}
	params := make(map[string]interface{})
	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Printf("bbf#back#verify#marshal_error, err: %v", err)
		return err
	}
	err = json.Unmarshal(jsonData, &params)
	if err != nil {
		log.Printf("bbf#back#verify#unmarshal_error, err: %v", err)
		return err
	}

	// Verify signature
	flag := utils.Verify(params, key)
	if !flag {
		reqJson, _ := json.Marshal(req)
		log.Printf("bbf#back#verify#fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}
	
	//开始处理
	return processor(req)
}
