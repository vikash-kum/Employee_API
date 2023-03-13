// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"employee/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/employee",
			beego.NSInclude(&controllers.EmployeeController{}),
		),
		beego.NSNamespace("/address",
			beego.NSInclude(&controllers.AddressController{}),
		),
		beego.NSNamespace("/origination",
			beego.NSInclude(&controllers.OriginationController{}),
		),
	)
	beego.AddNamespace(ns)
}
