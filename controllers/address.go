package controllers

import (
	"employee/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"strconv"
)

// AddressController Operations about address
type AddressController struct {
	beego.Controller
}

// Post @Title Post
// @Description  create user
// @Param body body models.Address true "The address to add"
// @Success	200	{int} models.Address.Id
// @Failure 403 body is empty
// @router / [post]
func (a *AddressController) Post() {
	var address models.Address
	json.Unmarshal(a.Ctx.Input.RequestBody, &address)
	// Custom validation
	valid := validation.Validation{}
	b, err := valid.Valid(&address)

	fmt.Println(address.Employee.EmpId)
	if err != nil {
		log.Println(err.Error())
	}
	if !b {
		a.Data["json"] = valid.Errors
	} else {
		result := models.AddAddress(&address)
		a.Data["json"] = result
	}
	a.ServeJSON()
}

// Get @Title Get
// @Description get all addresses
// @Success 200 {object} models.Address
// @Failure 403 Address is empty
// @router / [get]
func (a *AddressController) Get() {
	addresses, err := models.GetAllAddress()
	if err != nil {
		a.Data["json"] = err.Error()
	}
	a.Data["json"] = addresses
	a.ServeJSON()
}

// GetAddressById @Title Get
// @Description find address by addressId
// @Param aid path string true  "the addressed you want to get"
// @Success 200 {object} models.Address
// @Failure 403 addressId is empty
// @router /:aid [get]
func (a *AddressController) GetAddressById() {
	id := a.Ctx.Input.Param(":aid")
	aid, _ := strconv.Atoi(id)
	addresses, err := models.GetAddressById(aid)
	if err != nil {
		a.Data["json"] = err
	}
	a.Data["json"] = addresses
	a.ServeJSON()
}

// Put @Title Put
// @Description update the address
// @Param	aid	 path 		string		true	"The addressId you want to update"
// @Param	body body models.Address true "the addressId you want to update"
// @Success 200 {object} models.Address
// @Failure 403 addressId is empty
// @router /:aid [put]
func (a *AddressController) Put() {
	id := a.Ctx.Input.Param(":aid")
	aid, _ := strconv.Atoi(id)
	if aid > 0 {
		var address models.Address
		json.Unmarshal(a.Ctx.Input.RequestBody, &address)
		result := models.UpdateAddress(aid, &address)
		a.Data["json"] = result
	}
	a.ServeJSON()
}

// Delete @Title Delete
// @Description delete the object
// @Param	aid		path 	string	true		"The addressId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 addressId is empty
// @router /:aid [delete]
func (a *AddressController) Delete() {
	id := a.Ctx.Input.Param(":aid")
	aid, _ := strconv.Atoi(id)
	a.Data["json"] = models.DeleteAddressById(aid)
	a.ServeJSON()
}
