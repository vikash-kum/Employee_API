package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	beego.GlobalControllerRouter["employee/controllers:OriginationController"] = append(beego.GlobalControllerRouter["employee/controllers:OriginationController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:OriginationController"] = append(beego.GlobalControllerRouter["employee/controllers:OriginationController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:OriginationController"] = append(beego.GlobalControllerRouter["employee/controllers:OriginationController"],
		beego.ControllerComments{
			Method:           "GetOriginationById",
			Router:           "/:oid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:OriginationController"] = append(beego.GlobalControllerRouter["employee/controllers:OriginationController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:oid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:OriginationController"] = append(beego.GlobalControllerRouter["employee/controllers:OriginationController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:oid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:AddressController"] = append(beego.GlobalControllerRouter["employee/controllers:AddressController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:AddressController"] = append(beego.GlobalControllerRouter["employee/controllers:AddressController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:AddressController"] = append(beego.GlobalControllerRouter["employee/controllers:AddressController"],
		beego.ControllerComments{
			Method:           "GetAddressById",
			Router:           "/:aid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:AddressController"] = append(beego.GlobalControllerRouter["employee/controllers:AddressController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:aid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:AddressController"] = append(beego.GlobalControllerRouter["employee/controllers:AddressController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:aid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// Employee Controller
	beego.GlobalControllerRouter["employee/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["employee/controllers:EmployeeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["employee/controllers:EmployeeController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["employee/controllers:EmployeeController"],
		beego.ControllerComments{
			Method:           "GetEmployeeById",
			Router:           "/:eid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["employee/controllers:EmployeeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:eid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["employee/controllers:EmployeeController"] = append(beego.GlobalControllerRouter["employee/controllers:EmployeeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:eid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
