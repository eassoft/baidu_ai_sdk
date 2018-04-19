package main

import (
	//"github.com/eassoft/baidu_ai_sdk/config"
	//"github.com/eassoft/baidu_ai_sdk/routers"

	"hm/eclibs/logger"

	"os"
	"os/signal"
	"syscall"
	//"time"

	//"github.com/valyala/fasthttp"

	//"golang.org/x/crypto/acme/autocert"
	"github.com/eassoft/baidu_ai_sdk/vision/face"
	//"github.com/eassoft/baidu_ai_sdk/vision/ocr"
)

func main() {

	logger.Debug("baidu_ai_sdk Ver:1.0.0")

	//LoadConf()

	defer func() { // 必須要先声明defer,否则不能捕获到panic异常
		if err := recover(); err != nil {
			logger.Error("main ", err) // err其实就是panic异常的信息
		}
	}()
	signalChan := make(chan os.Signal, 1)                      //创建一个信号量的chan，缓存为1，（0,1）意义不大 让进程优雅的退出
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) //让进程收集信号量。
	//图片比较
	/*pic1 := "pics/1.jpg"
		pic2 := "pics/4.png"
		face.Match(pic1, pic2)
	      //图片验证
		pic1 := "pics/p2.jpg"
		face.DetectAndAnalysis(pic1)
	    //读取身份证
		idcard := "pics/身份证正面.JPG"
		ocr.IdCardRecognize(idcard)
	*/
	/*
		// 在线活体认证
		pic1 := "pics/身份证正面.jpg"
		face.FaceVerify(pic1)
	*/

	// 在线活体认证
	uid := "123456789"
	group_id := "rckjstamp"

	pic1 := "pics/4.png"
	/*user_info := "liuxc"
	face.UserAdd(pic1, uid, group_id, user_info)
	*/
	face.UserVerify(pic1, uid, group_id)

}
