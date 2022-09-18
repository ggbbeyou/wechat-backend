package common

import (
	"reflect"
)

/*
* @Author: chuang
* @Date:   2022/9/10 17:53
 */

// ObjCopy 将obj1中的属性值拷贝到obj2
func ObjCopy(src, dst interface{}) {
	st := reflect.TypeOf(src)
	sv := reflect.ValueOf(src)
	dt := reflect.TypeOf(dst)
	dv := reflect.ValueOf(dst)
	if st.Kind() == reflect.Ptr { //处理指针
		st = st.Elem()
		sv = sv.Elem()
	}
	if dt.Kind() == reflect.Ptr { //处理指针
		dt = dt.Elem()
	}
	if st.Kind() != reflect.Struct || dt.Kind() != reflect.Struct { //如果不是struct类型，直接返回dst
		return
	}

	dv = reflect.ValueOf(dv.Interface())
	// 遍历TypeOf 类型
	for i := 0; i < dt.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
		f := dt.Field(i) //通过这个i作为它的索引，从0开始来取得它的字段
		dVal := dv.Elem().Field(i)
		sVal := sv.FieldByName(f.Name)
		//fmt.Println(dVal.CanSet())
		//src数据有效，且dst字段能赋值,类型一致
		if sVal.IsValid() && dVal.CanSet() && f.Type.Kind() == sVal.Type().Kind() {
			dVal.Set(sVal)
		}
	}
	//return dst
}
