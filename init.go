package main

import (
	"hm/eclibs/logger"
)

func init() {

	initlog()
	//initLevelgo()
}

func initlog() {
	logger.SetRollingDaily("logs", "log.log") //设置日记打印文件夹和日记信息的输入文件

	logger.SetConsole(true) //设置为控制台输出

	logger.SetLevel(1) //设置错误级别为debug
}
