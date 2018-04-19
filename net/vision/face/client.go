package face

import (
	sdk "github.com/eassoft/baidu_ai_sdk/net"
)

type FaceClient struct {
	*sdk.Client
}

func NewFaceClient(key, secret string) *FaceClient {
	return &FaceClient{
		Client: sdk.NewClient(key, secret),
	}
}
