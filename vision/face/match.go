package face

import (
	"fmt"

	"github.com/eassoft/baidu_ai_sdk/config"
	"github.com/eassoft/baidu_ai_sdk/net/vision"
	"github.com/eassoft/baidu_ai_sdk/net/vision/face"
)

//比较两个图片的相似度
func Match(pic1, pic2 string) {
	client := face.NewFaceClient(config.B_Api_Key, config.B_Secret_Key)

	rs, err := client.Match(
		vision.MustFromFile(pic1),
		vision.MustFromFile(pic2),
		map[string]interface{}{
			"ext_fields":     "qualities",                 //返回质量信息，取值固定，目前支持qualities(质量检测)(对所有图片都会做改处理)
			"image_liveness": "faceliveness,faceliveness", //返回的活体信息，“faceliveness,faceliveness” 表示对比对的两张图片都做活体检测；“,faceliveness” 表示对第一张图片不做活体检测、第二张图做活体检测；“faceliveness,” 表示对第一张图片做活体检测、第二张图不做活体检测；
			"types":          "7,7",
		},
	)

	if err != nil {
		panic(err)
	}
	fmt.Println(rs.Dump())
}
