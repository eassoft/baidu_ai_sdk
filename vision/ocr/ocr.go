package ocr

import (
	"fmt"

	"github.com/easycode/baidu_ai_sdk/config"
	"github.com/easycode/baidu_ai_sdk/net/vision"
	"github.com/easycode/baidu_ai_sdk/net/vision/ocr"
)

var client *ocr.OCRClient

func init() {
	client = ocr.NewOCRClient(config.B_Api_Key, config.B_Secret_Key)
}

func main() {
	GeneralRecognizeBasic()
	GeneralRecognizeEnhanced()
}

func GeneralRecognizeBasic() {
	rs, err := client.GeneralRecognizeBasic(
		vision.MustFromFile("ocr.jpg"),
		ocr.DetectDirection(),       //是否检测图像朝向，默认不检测
		ocr.DetectLanguage(),        //是否检测语言，默认不检测。
		ocr.LanguageType("CHN_ENG"), //识别语言类型，默认为CHN_ENG。
		ocr.WithProbability(),       //是否返回识别结果中每一行的置信度
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(rs.ToString())
}

//一般性识别
func GeneralRecognizeEnhanced() {

	resp, err := client.GeneralRecognizeEnhanced(
		vision.MustFromFile("ocr.jpg"),
		ocr.DetectDirection(),
		ocr.DetectLanguage(),
		ocr.LanguageType("CHN_ENG"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ToString())

}

//身份证识别
func IdCardRecognize(idcard string) {

	resp, err := client.IdCardRecognize(
		vision.MustFromFile(idcard),
		ocr.DetectDirection(),
		ocr.DetectLanguage(),
		ocr.LanguageType("CHN_ENG"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ToString())

}
