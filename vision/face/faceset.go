package face

import (
	"fmt"

	"github.com/eassoft/baidu_ai_sdk/config"
	"github.com/eassoft/baidu_ai_sdk/net/vision"
	"github.com/eassoft/baidu_ai_sdk/net/vision/face"
)

//增加图片图库
func UserAdd(picpath, uid, group_id, user_info string) {
	client := face.NewFaceClient(config.B_Api_Key, config.B_Secret_Key)
	options := map[string]interface{}{
		"uid":         uid, //uid	是	string	用户id（由数字、字母、下划线组成），长度限制128B。
		"group_id":    group_id,
		"user_info":   user_info, //	是	string	用户资料，长度限制256B",
		"action_type": "app",     //参数包含app、replace。如果为“replace”，则每次注册时进行替换replace（新增或更新）操作，默认为append操作。例如：uid在库中已经存在时，对此uid重复注册时，新注册的图片默认会追加到该uid下，如果手动选择action_type:replace，则会用新图替换库中该uid下所有图片。
	}
	rs, err := client.UserAdd(
		vision.MustFromFile(picpath),
		options,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.ToString())
}

//修改图片
func UserUpdate(picpath, uid, group_id, user_info string) {
	client := face.NewFaceClient(config.B_Api_Key, config.B_Secret_Key)
	options := map[string]interface{}{
		"uid":         uid, //uid	是	string	用户id（由数字、字母、下划线组成），长度限制128B。
		"group_id":    group_id,
		"user_info":   user_info, //	是	string	用户资料，长度限制256B",
		"action_type": "replace", //参数包含app、replace。如果为“replace”，则每次注册时进行替换replace（新增或更新）操作，默认为append操作。例如：uid在库中已经存在时，对此uid重复注册时，新注册的图片默认会追加到该uid下，如果手动选择action_type:replace，则会用新图替换库中该uid下所有图片。
	}
	rs, err := client.UserAdd(
		vision.MustFromFile(picpath),
		options,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.ToString())
}

/*
删除图片
*/
func UserDelete(uid, group_id string) {
	client := face.NewFaceClient(config.B_Api_Key, config.B_Secret_Key)
	options := map[string]interface{}{
		"uid":      uid, //uid	是	string	用户id（由数字、字母、下划线组成），长度限制128B。
		"group_id": group_id,
	}
	rs, err := client.UserDelete(
		options,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.ToString())
}

//对比图片
func UserVerify(picpath, uid, group_id string) {
	client := face.NewFaceClient(config.B_Api_Key, config.B_Secret_Key)
	options := map[string]interface{}{
		"uid":      uid, //uid	是	string	用户id（由数字、字母、下划线组成），长度限制128B。
		"group_id": group_id,
		"top_num":  1, //	uint32	返回匹配得分top数，默认为1
		//"ext_fields": "replace", //string	特殊返回信息，多个用逗号分隔，取值固定: 目前支持faceliveness(活体检测)。注：需要用于判断活体的图片，图片中的人脸像素面积需要不小于100px*100px，人脸长宽与图片长宽比例，不小于1/3
	}

	rs, err := client.UserVerify(
		vision.MustFromFile(picpath),
		options,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.ToString())
}
