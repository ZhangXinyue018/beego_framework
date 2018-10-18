package controllers

import (
	"beego_framework/bean"
	"beego_framework/domain"
	"beego_framework/service"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
)

type TempController struct {
	MainController
	TestService service.ITestService
}

// @Title Test
// @Description test
// @Param	str	query string	false	"string for qrcode"
// @Success 200 {string}
// @router /test [get]
func (testController *TempController) Test() {
	defer testController.HandleError()
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
	defer testController.HandleError()
	//testOneAsync := make(chan commons.Async)
	//defer commons.ClearGoRoutine(testOneAsync)
	//go func() {
	//	defer commons.DeferErrorAsync(testOneAsync)
	//	testOneAsync <- commons.Async{Error: fmt.Errorf("test")}
	//}()
	//
	//testTwoAsync := make(chan commons.Async)
	//defer commons.ClearGoRoutine(testTwoAsync)
	//go func() {
	//	defer commons.DeferErrorAsync(testTwoAsync)
	//	testTwoAsync <- commons.Async{Result: 2}
	//}()
	//fmt.Println(commons.GetAsyncResult(testOneAsync).(int))
	//fmt.Println(commons.GetAsyncResult(testTwoAsync).(int))
	fmt.Println("---------------------------------")
	fmt.Println(testController.TestService.Test())
	var testmysql domain.TestMysql
	json.Unmarshal([]byte("{\"id\": 13, \"test\": \"hello\"}"), &testmysql)
	fmt.Printf("%v", testmysql)
	bean.ExchangerServiceBean.UpdateExchangerRate()
	fmt.Println("---------------------------------")
	testController.Data["json"] = "ok"
	testController.ServeJSON()
}
