package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"strconv"
)

func init() {
	orm.RegisterModel(new(Origination))
}

type Origination struct {
	Id              int       `orm:"auto"`
	Designation     string    `orm:"column(designation);size(50)";valid:"Required";json:"designation"`
	OriginationName string    `orm:"column(origination_name);size(500)";valid:"Required";json:"originationName"`
	DurationInYears int       `orm:"column(duration)";valid:"Required";json:"duration"`
	Employee        *Employee `orm:"rel(fk)"`
}

func (o *Origination) Valid(v *validation.Validation) {
	if o.Designation == "" {
		v.SetError("Designation", "Field Should not be empty")
	}
	if o.DurationInYears < 0 {
		v.SetError("Duration", "Should not be negative")
	}
	if o.OriginationName == "" {
		v.SetError("OriginationName", "Should not be empty")
	}
}

func GetOriginationById(id int) (*Origination, error) {
	o := orm.NewOrm()
	origination := &Origination{Id: id}
	err := o.Read(origination, "Id")
	origination.Employee = nil
	if err != nil {
		return nil, err
	} else {
		return origination, nil
	}
}

func GetAllOrigination() ([]*Origination, error) {
	o := orm.NewOrm()
	var originations []*Origination
	_, err := o.QueryTable("origination").All(&originations)
	totalOriginationLength := len(originations)
	for i := 0; i < totalOriginationLength; i++ {
		originations[i].Employee = nil
	}
	if err != nil {
		return nil, err
	} else {
		return originations, nil
	}

}

func AddOrigination(origination *Origination) string {
	o := orm.NewOrm()
	id, err := o.Insert(origination)
	if err != nil {
		return err.Error()
	} else {
		return "Successfully origination added and origination Id is: " + strconv.FormatInt(id, 10)
	}
}

func DeleteOrigination(id int) string {
	o := orm.NewOrm()
	origination := &Origination{Id: id}
	oid, err := o.Delete(origination)
	if err != nil {
		return err.Error()
	} else {
		return "Origination is deleted with ID: " + strconv.FormatInt(oid, 10)
	}
}

func UpdateOrigination(id int, origination *Origination) string {
	o := orm.NewOrm()
	origination.Id = id
	originationId, err := o.Update(&origination)
	if err != nil {
		return err.Error()
	} else {
		return "User delete successfully with id: " + strconv.FormatInt(originationId, 10)
	}
}

func GetOriginationOfUser(uid int) ([]*Origination, error) {
	o := orm.NewOrm()
	var employeeOrigination []*Origination
	qs := o.QueryTable("origination").Filter("employee__EmpId", uid)
	_, err := qs.All(&employeeOrigination)
	totalAddress := len(employeeOrigination)
	for i := 0; i < totalAddress; i++ {
		employeeOrigination[i].Employee = nil
	}
	if err != nil {
		return nil, err
	} else {
		return employeeOrigination, nil
	}
}
