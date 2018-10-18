package main

import (
	_ "beego_framework/bean"
	_ "beego_framework/routers"
	"net/http"
	_ "net/http/pprof"

	"beego_framework/cronjobs"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
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
