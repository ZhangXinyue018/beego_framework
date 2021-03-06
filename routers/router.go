// @APIVersion 1.0.0
// @Title test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego_framework/controllers"

	"github.com/astaxie/beego"
	"beego_framework/bean"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/version",
			beego.NSInclude(
				&controllers.VersionController{},
			),
		),
		beego.NSNamespace("/ws",
			beego.NSInclude(
				&controllers.WebSocketController{},
			),
		),
		beego.NSNamespace("/temp",
			beego.NSInclude(
				&controllers.TempController{
					TestService: bean.TestServiceBean,
				},
			),
		),
	)
	beego.AddNamespace(ns)
}
