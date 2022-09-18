package common

/*
* @Author: chuang
* @Date:   2022/9/17 8:34
 */

//redis key

const (
	MSGCACHE = "msgcache:" //离线消息缓存key

	CHATLIST = "chatlist:" //聊天列表

	CHATLIST_EXPIRE = 3600 * 2
)
