package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/session/mysql"
	_"github.com/go-sql-driver/mysql"
)

type China_division struct {
	Id          int64
	Name        string `orm:"size(128)"`
	Code        string
	ParentCode  int
	CoordinateY float64
	CoordinateX float64
	CreatedAt   int
	UpdatedAt   int
}

type China_division_join struct {
	ProvideCode    string
	CityCode       string
}

func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/api_stag?charset=utf8", 30)
	orm.RegisterModel(new(China_division))
}

// AddChina_division insert a new China_division into database and returns
// last inserted Id on success.
func AddChina_division(m *China_division) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChina_divisionById retrieves China_division by Id. Returns error if
// Id doesn't exist
func GetChina_divisionById(id int64) (v *China_division, err error) {
	o := orm.NewOrm()
	v = &China_division{Id: id}
	if err = o.QueryTable(new(China_division)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

//根据code返回数据
func GetChina_divisionByCode(code string) (v *China_division, err error) {
	o := orm.NewOrm()
	v = &China_division{Code: code}
	if err = o.QueryTable(new(China_division)).Filter("Code", code).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChina_division retrieves all China_division matches certain condition. Returns empty list if
// no records exist
func GetAllChina_division(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(China_division))
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

	var l []China_division
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

// UpdateChina_division updates China_division by Id and returns error if
// the record to be updated doesn't exist
func UpdateChina_divisionById(m *China_division) (err error) {
	o := orm.NewOrm()
	v := China_division{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChina_division deletes China_division by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChina_division(id int64) (err error) {
	o := orm.NewOrm()
	v := China_division{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&China_division{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//根据当前的code获取上面两层CODE
func FindParentsCodeByAdcode(code string) (result China_division_join){
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.Raw("SELECT `provide`.`code` AS `provideCode`, `city`.`code` AS `cityCode` FROM `china_division` INNER JOIN " +
		"`china_division` `city` ON city.code = china_division.parent_code INNER JOIN `china_division` `provide` ON city.parent_code = provide.code " +
		"WHERE china_division.code = ?",code).Values(&maps)
	if err != nil {
		fmt.Println(fmt.Sprintf("查询china_division: 出错:%v", err))
	}
	//fmt.Println(maps)

	if len(maps) == 0 {
		return China_division_join{}
	}
	provideCode := maps[0]["provideCode"]
	cityCode := maps[0]["cityCode"]

	provide, ok := provideCode.(string)
	if !ok {
		fmt.Println("断言provideCode失败")
	}
	city, ok := cityCode.(string)
	if !ok {
		fmt.Println("断言cityCode失败")
	}
	return China_division_join{provide,city}
}
