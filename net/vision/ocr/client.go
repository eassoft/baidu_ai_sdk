package ocr

import "github.com/eassoft/baidu_ai_sdk/net"

type OCRClient struct {
	*net.Client
}

func NewOCRClient(apiKey, secretKey string) *OCRClient {
	return &OCRClient{
		Client: net.NewClient(apiKey, secretKey),
	}
}
