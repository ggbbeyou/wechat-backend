package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

/*
* @Author: chuang
* @Date:   2023/2/5 9:19
 */

func Str2Int(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		logx.Error(err)
		panic(0)
	}
	return res
}

func Str2Int64(str string) int64 {
	res, err := strconv.Atoi(str)
	if err != nil {
		logx.Error(err)
		panic(0)
	}
	return int64(res)
}
