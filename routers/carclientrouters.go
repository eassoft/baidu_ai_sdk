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
	/***********xber beg***************/
	orderctrl      = new(controllers.OrderCtrl)
	membercarsctrl = new(controllers.MemberCarsCtrl)

	/***********xber end***************/
	washsitectrl = new(controllers.WashSiteCtrl)

	usercenterctrl = new(controllers.UserCenterCtrl)

	couponsctrl      = new(controllers.CouponsCtrl)
	moneyChangeCtrl  = new(controllers.MoneyChangeCtrl)
	pointsChangeCtrl = new(controllers.PointsChangeCtrl)

	carBrandCtrl    = new(controllers.CarBrandCtrl)
	suggestionsCtrl = new(controllers.SuggestionsCtrl)
	wxuserinfoCtrl  = new(controllers.WxUserInfoCtrl)
)

func RoutersClient(e *echo.Echo) *echo.Echo {
	logger.Debug("OrderRouters")

	e.GET("/order/:uid", orderctrl.Get)
	e.GET("/order/q/:method", orderctrl.Query)
	e.POST("/order", orderctrl.Post)
	e.PUT("/order/:uid", orderctrl.Put)
	e.POST("/order/:method", orderctrl.PATCH)

	//我的爱车
	e.GET("/membercars/:uid", membercarsctrl.Get)
	e.GET("/membercars/q/:method", membercarsctrl.Query)
	e.POST("/membercars", membercarsctrl.Post)
	e.PUT("/membercars/:uid", membercarsctrl.Put)
	e.POST("/membercars/:method", membercarsctrl.Patch)

	logger.Debug("WashApiRouters")
	//站点
	e.GET("/washsite/:uid", washsitectrl.Get)
	e.GET("/washsite/q/:method", washsitectrl.Query)
	e.POST("/washsite", washsitectrl.Post)
	e.PUT("/washsite", washsitectrl.Put)
	//e.PATCH("/washsite/:method", washsitectrl.PATCH)
	//订单

	//用户中心
	//e.GET("/usercenter/:uid", usercenterctrl.Get)
	e.GET("/usercenter/:method", usercenterctrl.Query)
	e.POST("/usercenter", usercenterctrl.Post)
	//e.PUT("/usercenter/:uid", usercenterctrl.Put)
	e.POST("/usercenter/:method", usercenterctrl.PATCH)
	//

	//我的优惠券
	e.GET("/coupons/:uid", couponsctrl.Get)
	e.GET("/coupons/q/:method", couponsctrl.Query)
	e.POST("/coupons", couponsctrl.Post)
	e.PUT("/coupons", couponsctrl.Put)

	//充值
	e.GET("/money/:uid", moneyChangeCtrl.Get)
	e.GET("/money/q/:method", moneyChangeCtrl.Query)
	e.POST("/money", moneyChangeCtrl.Post)
	e.PUT("/money", moneyChangeCtrl.Put)

	//积分
	e.GET("/points/:uid", pointsChangeCtrl.Get)
	e.GET("/points/q/:method", pointsChangeCtrl.Query)
	e.POST("/points", pointsChangeCtrl.Post)
	e.PUT("/points", pointsChangeCtrl.Put)

	//品牌型号
	//e.GET("/brand/:uid", carBrandCtrl.Get)
	e.GET("/brand/q/:method", carBrandCtrl.Query)
	e.POST("/brand", carBrandCtrl.Post)
	e.PUT("/brand", carBrandCtrl.Put)

	//意见建议
	e.GET("/suggestions/:uid", suggestionsCtrl.Get)
	e.GET("/suggestions/q/:method", suggestionsCtrl.Query)
	e.POST("/suggestions", suggestionsCtrl.Post)
	e.PUT("/suggestions/:uid", suggestionsCtrl.Put)
	e.DELETE("/suggestions/:uid", suggestionsCtrl.Delete)

	/*站点套餐*/
	wash2saleductCtrl := new(controllers.Wash2SaleDuctCtrl)
	e.GET("/wash2saleduct/:uid", wash2saleductCtrl.Get)
	e.GET("/wash2saleduct/q/:method", wash2saleductCtrl.Query)
	e.POST("/wash2saleduct", wash2saleductCtrl.Post)
	e.POST("/wash2saleduct/:method", wash2saleductCtrl.Patch)
	e.PUT("/wash2saleduct/:uid", wash2saleductCtrl.Put)
	e.DELETE("/wash2saleduct/:uid", wash2saleductCtrl.Delete)
	/*微信用户*/
	wxuserinfoCtrl = new(controllers.WxUserInfoCtrl)
	e.GET("/wxuserinfo/:uid", wxuserinfoCtrl.Get)
	e.GET("/wxuserinfo/q/:method", wxuserinfoCtrl.Query)
	e.POST("/wxuserinfo", wxuserinfoCtrl.Post)
	e.PUT("/wxuserinfo/:uid", wxuserinfoCtrl.Put)
	e.POST("/wxuserinfo/:method", wxuserinfoCtrl.Patch)
	e.DELETE("/wxuserinfo/:uid", wxuserinfoCtrl.Delete)

	return e

}
