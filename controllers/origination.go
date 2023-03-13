package controllers

import (
	"employee/models"
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"strconv"
)

// OriginationController Operations about address
type OriginationController struct {
	beego.Controller
}

// Post @Title Post
// @Description  create origination
// @Param body body models.Origination true "The origination to add"
// @Success	200	{int} models.Origination.Id
// @Failure 403 body is empty
// @router / [post]
func (o *OriginationController) Post() {
	var origination models.Origination
	json.Unmarshal(o.Ctx.Input.RequestBody, &origination)
	// Custom validation
	valid := validation.Validation{}
	b, err := valid.Valid(&origination)
	if err != nil {
		log.Println(err.Error())
	}
	if !b {
		o.Data["json"] = valid.Errors
	} else {
		result := models.AddOrigination(&origination)
		o.Data["json"] = result
	}
	o.ServeJSON()
}

// Get @Title Get
// @Description get all origination
// @Success 200 {object} models.Origination
// @Failure 403 origination is empty
// @router / [get]
func (o *OriginationController) Get() {
	origination, err := models.GetAllOrigination()
	if err != nil {
		o.Data["json"] = err.Error()
	}
	o.Data["json"] = origination
	o.ServeJSON()
}

// @Title Get
// @Description find origination by originationId
// @Param oid path string true  "the origination you want to get"
// @Success 200 {object} models.Origination
// @Failure 403 originationId is empty
// @router /:oid [get]
func (o *OriginationController) GetOriginationById() {
	uid := o.Ctx.Input.Param(":oid")
	id, _ := strconv.Atoi(uid)
	origination, err := models.GetOriginationById(id)
	if err != nil {
		o.Data["json"] = err
	}
	o.Data["json"] = origination
	o.ServeJSON()
}

// Put @Title Put
// @Description update the origination
// @Param	oid	 path 		string		true	"The originationId you want to update"
// @Param	body body models.Origination true "the originationId you want to update"
// @Success 200 {object} models.Origination
// @Failure 403 originationId is empty
// @router /:oid [put]
func (o *OriginationController) Put() {
	id := o.Ctx.Input.Param(":oid")
	aid, _ := strconv.Atoi(id)
	if aid > 0 {
		var origination models.Origination
		json.Unmarshal(o.Ctx.Input.RequestBody, &origination)
		result := models.UpdateOrigination(aid, &origination)
		o.Data["json"] = result
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	oid		path 	string	true		"The originationId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 originationId is empty
// @router /:oid [delete]
func (o *OriginationController) Delete() {
	id := o.Ctx.Input.Param(":oid")
	aid, _ := strconv.Atoi(id)
	o.Data["json"] = models.DeleteOrigination(aid)
	o.ServeJSON()
}
