package controllers

import (
	"github.com/astaxie/beego"
	"beego_framework/bean"
	)

type MainController struct {
	beego.Controller
}

// @Title GetErrorMap
// @Description get error map
// @Success 200 {object} response.BaseResp
// @router / [get]
func (u *MainController) Get() {
	defer u.HandleError()
	resultMap := make(map[string]string, len(bean.ErrorMap))

	for _, value := range bean.ErrorMap {
		resultMap[value.ErrorCode] = value.ErrorMessage
	}
	u.Data["json"] = resultMap
	u.ServeJSON()
}

func (u *MainController) HandleError()(){
	if x := recover(); x != nil {
		u.Data["json"] = bean.HandleError(x)
		u.ServeJSON()
	}
}