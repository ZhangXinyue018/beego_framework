package controllers

import (
	"github.com/skip2/go-qrcode"
	"fmt"
	)

type TempController struct {
	MainController
}

// @Title Test
// @Description test
// @Param	str	query string	false	"string for qrcode"
// @Success 200 {string}
// @router /test [get]
func (testController *TempController) Test() {
	testController.HandleError()
	str := testController.GetString("str")
	result, err := qrcode.Encode(str, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("write error")
	}

	testController.Ctx.Output.ContentType("jpeg")
	testController.Ctx.Output.Body(result)

	//testController.Data["json"] = "ok"
	//testController.ServeJSON()
}

// @Title Get
// @Description test
// @Success 200 {string}
// @router / [get]
func (testController *TempController) Get() {
	testController.HandleError()
	testController.Data["json"] = "ok"
	testController.ServeJSON()
}
