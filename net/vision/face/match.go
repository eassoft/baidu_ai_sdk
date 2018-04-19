package face

import (
	"fmt"

	"github.com/eassoft/baidu_ai_sdk/net/vision"
	"github.com/imroc/req"
)

const (
	FACE_MATCH_URL = "https://aip.baidubce.com/rest/2.0/face/v2/match"
)

//比较两个图片
func (fc *FaceClient) Match(img1, img2 *vision.Image, options map[string]interface{}) (*FaceResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}

	imgContent1, err := img1.Base64Encode()
	imgContent2, err := img2.Base64Encode()
	if err != nil { //任意一个出错都不行
		return nil, err
	}

	options["images"] = fmt.Sprintf("%s,%s", imgContent1, imgContent2)

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	url := FACE_MATCH_URL + "?access_token=" + fc.AccessToken
	resp, err := req.Post(url, header, req.Param(options))
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}
