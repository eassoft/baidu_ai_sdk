package routers

import (

	//"hm/imageserver/filters"

	"hm/eclibs/logger"
	"hm/flowerapi/controllers"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"

	//"github.com/valyala/fasthttp"
)

// 业务处理
var (
	wxapictrl = new(controllers.WxApiCtrl)
)

func WxApiRouters(e *echo.Echo) *echo.Echo {
	logger.Debug("WxApiRouters")

	e.GET("/wx/receiver", wxapictrl.VerifToken)
	e.POST("/wx/receiver", wxapictrl.Receiver)

	e.GET("/wx/api/:method", wxapictrl.Query)
	e.POST("/wx/api/:method", wxapictrl.Patch)

	e.POST("/wx/pay/jsapipay", wxapictrl.JSAPIPay)
	e.Any("/wx/pay/notify", wxapictrl.Notify)

	return e
}
