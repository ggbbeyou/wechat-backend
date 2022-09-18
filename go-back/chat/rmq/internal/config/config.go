package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
)

/*
* @Author: chuang
* @Date:   2022/9/15 11:10
 */

type Config struct {
	rest.RestConf
	KqConf kq.KqConf
}
