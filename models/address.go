package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"strconv"
)

func init() {
	orm.RegisterModel(new(Address))
}

// Address defines address struct
type Address struct {
	Id       int       `orm:"auto"`
	Address  string    `orm:"column(address);size(255)";valid:"Required";json:"address"`
	City     string    `orm:"column(city);size(20)";valid:"Required";json:"city"`
	Pin      string    `orm:"column(pin);size(6)";valid:"ZipCode";json:"pin"`
	Country  string    `orm:"column(country);size(50);default(india)";valid:"Required";json:"country"`
	Type     string    `orm:"column(type);size(50)";valid:"Required";json:"type"`
	Employee *Employee `orm:"rel(fk)";json:"employee"`
}

func (a *Address) Valid(v *validation.Validation) {
	if a.Address == "" {
		v.SetError("Address", "Should not be empty")
	}
	if a.Type == "" {
		v.SetError("Type", "Should be not empty")
	}
	if a.Country == "" {
		v.SetError("Country", "Should not be empty")
	}
	if a.City == "" {
		v.SetError("City", "City Should not be empty")
	}
}

func AddAddress(address *Address) string {
	o := orm.NewOrm()
	id, err := o.Insert(address)
	fmt.Println(address.Employee.EmpId)
	if err != nil {
		return err.Error()
	} else {
		return "Address added successfully and Id of address is: " + strconv.FormatInt(id, 10)
	}
}

func GetAddressOfUser(uid int) ([]*Address, error) {
	o := orm.NewOrm()
	var employeeAddress []*Address
	qs := o.QueryTable("address").Filter("employee__EmpId", uid)
	_, err := qs.All(&employeeAddress)

	totalAddress := len(employeeAddress)
	for i := 0; i < totalAddress; i++ {
		employeeAddress[i].Employee = nil
	}
	if err != nil {
		return nil, err
	} else {
		return employeeAddress, nil
	}
}

func GetAddressById(aid int) (*Address, error) {
	o := orm.NewOrm()
	address := &Address{Id: aid}
	err := o.Read(address, "Id")
	address.Employee = nil
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	} else {
		return address, nil
	}
}

func GetAllAddress() ([]Address, error) {
	o := orm.NewOrm()
	var addresses []Address
	_, err := o.QueryTable("address").All(&addresses)

	totalAddress := len(addresses)

	for i := 0; i < totalAddress; i++ {
		addresses[i].Employee = nil
	}
	if err != nil {
		return nil, err
	} else {
		return addresses, nil
	}

}

func DeleteAddressById(id int) string {
	o := orm.NewOrm()
	address := &Address{Id: id}
	aid, err := o.Delete(address)
	if err != nil {
		return err.Error()
	} else {
		return "Address of Id:" + strconv.FormatInt(aid, 10) + "is deleted successfully"
	}

}

func UpdateAddress(aid int, address *Address) string {
	o := orm.NewOrm()
	address.Id = aid
	id, err := o.Update(address)
	if err != nil {
		return err.Error()
	} else {
		return "address of Id: " + strconv.FormatInt(id, 10) + "is successfully updated"
	}
}
