package face

import (
	"fmt"

	"github.com/eassoft/baidu_ai_sdk/config"
	"github.com/eassoft/baidu_ai_sdk/net/vision"
	"github.com/eassoft/baidu_ai_sdk/net/vision/face"
)

//人脸鉴别 例如: 性别 年龄等
func DetectAndAnalysis(pic string) {
	client := face.NewFaceClient(config.B_Api_Key, config.B_Secret_Key)
	options := map[string]interface{}{
		"max_face_num": 10,
		"face_fields":  "age,beauty,expression,faceshape,gender,glasses,landmark,race,qualities",
	}
	rs, err := client.DetectAndAnalysis(
		vision.MustFromFile(pic),
		options,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.ToString())
}
