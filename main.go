package main

import (
	_ "beego_framework/routers"
	_ "beego_framework/bean"
	_ "net/http/pprof"

	"github.com/astaxie/beego"
	"fmt"
	"beego_framework/cronjobs"
	"beego_framework/bean"
	"net/http"
)

func main() {
	PerformSetUp()
	PerformSetUp()
	err := beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func PerformSetUp() {
	go cronjobs.StartCronjobs()
	go bean.WebSocketServiceBean.HandleChannelEvents()
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	//go func() {
	//	for {
	//		time.Sleep(3 * time.Second)
	//		fmt.Println(bean.WebSocketServiceBean.EventChannels)
	//		fmt.Println(bean.WebSocketServiceBean.ConnectionMap)
	//	}
	//}()
}