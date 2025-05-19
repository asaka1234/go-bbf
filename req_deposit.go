package go_bbf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-bbf/utils"
	"io/ioutil"
	"net/http"
)

// 下单
func (cli *Client) Deposit(req BbfDepositReq) (*BbfDepositRsp, error) {

	rawURL := cli.DepositURL

	// TODO 这里可以用mapstruct来简化
	// Prepare params map
	params := make(map[string]interface{})
	// Convert struct to map
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %v", err)
	}
	err = json.Unmarshal(jsonData, &params)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshal error: %v", err)
	}

	// Generate signature
	signature, err := utils.Sign(params, cli.AccessKey)
	if err != nil {
		return nil, fmt.Errorf("signature error: %v", err)
	}
	req.Signature = signature

	// Prepare request
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %v", err)
	}

	// Send HTTP request
	resp, err := http.Post(rawURL, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	// Parse response
	var result2 BbfDepositRsp
	err = json.Unmarshal(body, &result2)
	if err != nil {
		return nil, fmt.Errorf("parse response error: %v", err)
	}

	return &result2, nil
}
