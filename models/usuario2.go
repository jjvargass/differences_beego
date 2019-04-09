package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Usuario2 struct {
	Id       int
	Nombre   string `orm:"size(128)"`
	Apellido string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Usuario2))
}

// AddUsuario2 insert a new Usuario2 into database and returns
// last inserted Id on success.
func AddUsuario2(m *Usuario2) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsuario2ById retrieves Usuario2 by Id. Returns error if
// Id doesn't exist
func GetUsuario2ById(id int64) (v *Usuario2, err error) {
	o := orm.NewOrm()
	v = &Usuario2{Id: id}
	if err = o.QueryTable(new(Usuario2)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsuario2 retrieves all Usuario2 matches certain condition. Returns empty list if
// no records exist
func GetAllUsuario2(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Usuario2))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Usuario2
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUsuario2 updates Usuario2 by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsuario2ById(m *Usuario2) (err error) {
	o := orm.NewOrm()
	v := Usuario2{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsuario2 deletes Usuario2 by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsuario2(id int64) (err error) {
	o := orm.NewOrm()
	v := Usuario2{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Usuario2{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
