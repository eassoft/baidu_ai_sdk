package routers

import (

	//"hm/imageserver/filters"

	"hm/apiserver/controllers"
	"hm/eclibs/logger"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"

	//"github.com/valyala/fasthttp"
)

// 业务处理
var (
	stampctrl = new(controllers.StampCtrl)
)

func StampRouters(e *echo.Echo) *echo.Echo {
	logger.Debug("StampRouters")

	e.POST("/stamp/:pno", stampctrl.UploadHandler)
	//查看原圖
	e.GET("/stamp/:pno/img/:imgid", stampctrl.DownloadHandler)
	//缩略图
	e.GET("/stamp/:pno/thum/:imgid", stampctrl.DownloadThumHandler)
	//缩略图
	e.GET("/stamp/:pno/jpg/:imgid", stampctrl.DownloadJpgHandler)
	//动态图片服务
	e.GET("/stamp/:pno/autoimg/:width/:imgid", stampctrl.DownloadThumHandler)
	//生成二维码
	e.GET("/qrcode", stampctrl.QrCode)

	e.GET("/stamp/getstampimages", stampctrl.QueryAllImage)
	return e
}
