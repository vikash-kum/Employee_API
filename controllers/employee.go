package controllers

import (
	"employee/models"
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"strconv"
)

// EmployeeController Operations about employee
type EmployeeController struct {
	beego.Controller
}

// Post @Title	CreateEmployee
// @Description	create employee
// @Param	body body models.Employee	true "body for user content"
// @Success		200		{int}	models.Employee.Id
// @Failure		403		body	is	empty
// @router	/ [post]
func (u *EmployeeController) Post() {
	employee := models.Employee{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &employee)
	// Custom Validation
	valid := validation.Validation{}
	b, err := valid.Valid(&employee)
	if err != nil {
		log.Println(err.Error())
	}
	if !b {
		u.Data["json"] = valid.Errors
	} else {
		// If data is valid then save to the database
		result := models.AddEmployee(&employee)
		u.Data["json"] = result
	}
	u.ServeJSON()
}

// Get @Title Get All
// @Description get all employees
// @Success 200 {object} models.Employee
// @Failure 403  No employee data available
// @router / [get]
func (u *EmployeeController) Get() {
	employees, err := models.GetAllEmployee()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = employees
	}
	u.ServeJSON()
}

// GetEmployeeById @Title Get
// @Description get employee by id
// @Param  eid path string true "The employeeId you want to get"
// @Success 200 {object} models.Employee
// @Failure 403 employeeId not empty
// @router /:eid [get]
func (u *EmployeeController) GetEmployeeById() {
	uid := u.Ctx.Input.Param(":eid")
	id, _ := strconv.Atoi(uid)
	if id != 0 {
		employee, err := models.GetEmployeeById(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = employee
		}
	}
	u.ServeJSON()
}

// Put @Title Put
// @Description update the user
// @Param eid path string	 true "The employeeId you want to update"
// @Param body body models.Employee true 	"body for user content"
// @Success 200 {string} updated successfully
// @Failure 403 employeeId not empty
// @router /:eid [put]
func (u *EmployeeController) Put() {
	uid := u.Ctx.Input.Param(":eid")
	id, _ := strconv.Atoi(uid)
	if id > 0 {
		employee := models.Employee{}
		employee.EmpId = id
		json.Unmarshal(u.Ctx.Input.RequestBody, &employee)
		result := models.UpdateEmployee(id, &employee)
		u.Data["json"] = result
	} else {
		u.Data["json"] = "Please send correct ID"
	}
	u.ServeJSON()
}

// Delete @Title Delete
// @Description delete the object
// @Param	eid		path 	int	true "The userId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 eid is empty
// @router /:eid [delete]
func (u *EmployeeController) Delete() {
	uid := u.Ctx.Input.Param(":eid")
	id, _ := strconv.Atoi(uid)
	result := models.DeleteEmployee(id)
	u.Data["json"] = result
	u.ServeJSON()
}
