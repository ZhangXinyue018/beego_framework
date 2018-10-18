package factory

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	Logger *logs.BeeLogger
)

func GetLogger() (*logs.BeeLogger) {
	if Logger == nil {
		initLogger()
	}
	return Logger
}

func initLogger() () {
	Logger = getLog()
}

func getLog() (*logs.BeeLogger) {
	levelInt, err := beego.AppConfig.Int("loglevel")
	if err != nil {
		levelInt = 3
	}
	logger := logs.NewLogger(10000)
	//fileName := fmt.Sprintf("%d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	fileName := "app"
	fileStr := fmt.Sprintf(`{"filename": "log/%s.log"}`, fileName)
	logger.SetLogger(
		"file",
		fileStr)
	logger.EnableFuncCallDepth(true)
	logger.SetLevel(levelInt)
	logger.Async()
	return logger
}