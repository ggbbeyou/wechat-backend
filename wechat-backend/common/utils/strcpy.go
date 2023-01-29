package utils

import (
	"reflect"
	"strings"
	"time"
)

/*
* @Author: chuang
* @Date:   2023/1/10 14:29
 */

// StrCpy 将obj1中的属性值拷贝到obj2

func StrCpy(src, dst interface{}) {
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
		g, b := st.FieldByName(f.Name)
		if !b {
			continue
		}

		//src数据有效，且dst字段能赋值,类型一致
		if !sVal.IsValid() || !dVal.CanSet() {
			return
		}

		if f.Type.Kind() == g.Type.Kind() {
			dVal.Set(sVal)
			continue
		}

		//特殊判断 NullString、 NullInt64、 NullBool、NullTime
		if f.Type.Kind() == reflect.String && strings.Compare(g.Type.String(), "sql.NullString") == 0 {
			dVal.SetString(sVal.FieldByName("String").String())
		}

		if f.Type.Kind() == reflect.Int64 && (strings.Compare(g.Type.String(), "sql.NullInt64") == 0 || g.Type.Kind() == reflect.Uint64) {
			dVal.SetInt(sVal.FieldByName("Int64").Int())
		}

		if f.Type.Kind() == reflect.Bool && strings.Compare(g.Type.String(), "sql.NullBool") == 0 {
			dVal.SetBool(sVal.FieldByName("Bool").Bool())
		}

		if f.Type.Kind() == reflect.String && strings.Compare(g.Type.String(), "sql.NullTime") == 0 {
			time := sVal.FieldByName("Time").Interface().(time.Time)
			dVal.SetString(time.String())
		}

	}
}
