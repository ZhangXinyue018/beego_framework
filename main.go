package main

import (
	_ "beego_framework/routers"
	_ "beego_framework/bean"

	"github.com/astaxie/beego"
	"fmt"
	"beego_framework/cronjobs"
)

func main() {
	PerformCronjobs()
	err := beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func PerformCronjobs(){
	go cronjobs.CronUpdateExchangerRate()
}