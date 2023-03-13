package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"strconv"
	"strings"
)

func init() {
	orm.RegisterModel(new(Employee))
}

// Employee defines employee struct
type Employee struct {
	EmpId       int            `orm:"auto";json:"empId"`
	FirstName   string         `orm:"size(50)";valid:"Required";json:"firstName"`
	LastName    string         `orm:"size(50);null";json:"lastName"`
	Age         int            `orm:"null";valid:"Range(15, 130)";json:"age"`
	Email       string         `orm:"size(100);unique";valid:"Email";json:"email"`
	Password    string         `orm:"size(255)";valid:"Required";json:"password"`
	Mobile      string         `orm:"size(10);unique";valid:"Mobile";json:"mobile"`
	IsActive    bool           `orm:"default(false)";json:"-"`
	Address     []*Address     `orm:"reverse(many);null";json:"-"`
	Origination []*Origination `orm:"reverse(many);null";json:"-"`
}

func (u *Employee) Valid(v *validation.Validation) {
	if strings.Index(u.FirstName, "admin") != -1 || u.FirstName == "" {
		v.SetError("FirstName", "Can't contain 'admin' in Name and not be empty")
	}
}

func GetEmployeeById(id int) (*Employee, error) {
	o := orm.NewOrm()
	employee := &Employee{EmpId: id}
	err := error(nil)
	err = o.Read(employee, "EmpId")
	employee.Address, err = GetAddressOfUser(id)
	employee.Origination, _ = GetOriginationOfUser(id)
	if err != nil {
		return nil, err
	} else {
		return employee, nil
	}
}

func GetAllEmployee() ([]Employee, error) {
	o := orm.NewOrm()
	var employees []Employee
	_, err := o.QueryTable("employee").All(&employees)
	totalEmployee := len(employees)

	for i := 0; i < totalEmployee; i++ {
		employees[i].Address, _ = GetAddressOfUser(employees[i].EmpId)
		employees[i].Origination, _ = GetOriginationOfUser(employees[i].EmpId)
	}

	if err != nil {
		return nil, err
	} else {
		return employees, nil
	}

}

func AddEmployee(e *Employee) string {
	o := orm.NewOrm()
	id, err := o.Insert(e)
	if err != nil {
		return err.Error()
	} else {
		return "Successfully employee added and employee Id is: " + strconv.FormatInt(id, 10)
	}
}

func DeleteEmployee(id int) string {
	o := orm.NewOrm()
	employee := &Employee{EmpId: id}
	employeeId, err := o.Delete(employee)
	if err != nil {
		return err.Error()
	} else {
		return "Employee is deleted with ID: " + strconv.FormatInt(employeeId, 10)
	}
}

func UpdateEmployee(id int, employee *Employee) string {
	o := orm.NewOrm()
	employeeId, err := o.Update(&employee)
	if err != nil {
		return err.Error()
	} else {
		return "User delete successfully with id: " + strconv.FormatInt(employeeId, 10)
	}
}
