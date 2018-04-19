package face

import (
	"github.com/eassoft/baidu_ai_sdk/net/vision"
	"github.com/imroc/req"
)

const (
	FACE_USER_ADD_URL    = "https://aip.baidubce.com/rest/2.0/face/v2/faceset/user/add"
	FACE_USER_UPDATE_URL = "https://aip.baidubce.com/rest/2.0/face/v2/faceset/user/update"
	FACE_USER_DELETE_URL = "https://aip.baidubce.com/rest/2.0/face/v2/faceset/user/delete"
	//人脸登录 不需要uid 找到最相似的
	FACE_USER_IDENTIFY_URL = "https://aip.baidubce.com/rest/2.0/face/v2/identify"
	//验证人脸
	FACE_USER_VERIFY_URL = "https://aip.baidubce.com/rest/2.0/face/v2/verify"
)

// 人脸注册
func (fc *FaceClient) UserAdd(image *vision.Image, options map[string]interface{}) (*FaceResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}
	url := FACE_USER_ADD_URL + "?access_token=" + fc.AccessToken
	base64Str, err := image.Base64Encode()
	if err != nil {
		return nil, err
	}

	options["image"] = base64Str

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := req.Post(url, req.Param(options), header)
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}

// 人脸修改
func (fc *FaceClient) UserUpdate(image *vision.Image, options map[string]interface{}) (*FaceResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}
	url := FACE_USER_UPDATE_URL + "?access_token=" + fc.AccessToken
	base64Str, err := image.Base64Encode()
	if err != nil {
		return nil, err
	}

	options["image"] = base64Str

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := req.Post(url, req.Param(options), header)
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}

// 人脸delete
func (fc *FaceClient) UserDelete(options map[string]interface{}) (*FaceResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}
	url := FACE_USER_DELETE_URL + "?access_token=" + fc.AccessToken

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := req.Post(url, req.Param(options), header)
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}

// 人脸验证 人脸查找
func (fc *FaceClient) UserVerify(image *vision.Image, options map[string]interface{}) (*FaceResponse, error) {
	if err := fc.Auth(); err != nil {
		return nil, err
	}
	url := FACE_USER_VERIFY_URL + "?access_token=" + fc.AccessToken
	base64Str, err := image.Base64Encode()
	if err != nil {
		return nil, err
	}

	options["image"] = base64Str

	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := req.Post(url, req.Param(options), header)
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}
