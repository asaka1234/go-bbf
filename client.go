package go_bbf

import (
	"github.com/asaka1234/go-bbf/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID string // merchantId
	AccessKey  string // accessKey
	BackKey    string //

	DepositURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, accessKey, backKey, depositURL string) *Client {
	return &Client{
		MerchantID: merchantID,
		AccessKey:  accessKey,
		BackKey:    backKey,
		DepositURL: depositURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
