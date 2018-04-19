package routers

import (

	//"hm/imageserver/filters"

	"hm/apiserver/controllers"
	"hm/eclibs/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//"github.com/valyala/fasthttp"
)

// 业务处理
var (
	mdctrl = new(controllers.MDCtrl)
)

func Routers() *echo.Echo {
	logger.Debug("Routers")
	e := echo.New()
	//e.Use(middleware.BasicAuth())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	//e.GET("/md/:id", mdctrl.Get)
	e.GET("/md/:method", mdctrl.Query)
	e.POST("/md/:method", mdctrl.Post)

	smsctrl := new(controllers.SMSCtrl)
	e.GET("/sms/:method", smsctrl.Query)
	e.POST("/sms/:method", smsctrl.Post)
	/**添加AutoAPI**/
	autoapictrl := new(controllers.AutoApiCtrl)
	e.GET("/AutoApi/:module/:method", autoapictrl.Query)
	e.POST("/AutoApi/:module/:method", autoapictrl.Post)
	//e = RoutersClient(e)

	//e = RoutersWorker(e)
	//e = SiteManageRouters(e)
	//e = WxApiRouters(e)
	e = StampRouters(e)
	return e

}
