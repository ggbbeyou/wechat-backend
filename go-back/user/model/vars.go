package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

type BaseInfo struct {
	Uid       int64  `db:"uid"`      // 用户id 雪花算法
	Username  string `db:"username"` // 用户名
	Nickname  string `db:"nickname"` // 昵称
	Age       int64  `db:"age"`      // 年龄
	Avatar    string `db:"avatar"`   // 头像地址
	Introduce string `db:"introduce"`
	Email     string `db:"email"` // 邮箱
}

type FriendInfoDto struct {
	BaseInfo
}

type LoginUserInfo struct {
	BaseInfo
}

type GroupBaseInfo struct {
	Gid        int64  `db:"gid"`
	Uid        int64  `db:"uid"`
	Gname      string `db:"gname"`
	Avatar     string `db:"avatar"`
	Count      string `db:"count"`
	CreateTime string `db:"create_time"`
}

// UsersInfo 完整用户信息 个人设置修改可用
type UsersInfo struct {
	Id       int64  `db:"id"`       // 主键id
	Password string `db:"password"` // 密码 32位加密
	BaseInfo
	State       int64  `db:"state"`        // 0正常 1=禁用
	IsDelete    int64  `db:"is_delete"`    // 逻辑删除
	CreateTime  string `db:"create_time"`  // 创建时间
	UpdateTime  string `db:"update_time"`  // 更新信息时间
	LastTime    string `db:"last_time"`    // 上次登陆时间
	LastAddress string `db:"last_address"` // 上次登录地址
}

type GroupInfoDto struct {
	GroupBaseInfo
}

type GroupDetailDto struct {
	GroupBaseInfo
	Nickname string `db:"nickname"` // 昵称
}

// GroupIdsDto rpc方法 参数
type GidArr []int64