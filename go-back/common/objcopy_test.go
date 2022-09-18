package common

import (
	"fmt"
	"testing"
)

/*
* @Author: chuang
* @Date:   2022/9/10 18:01
 */

type User struct {
	A int
	B string
	C int
}

func TestAdd(t *testing.T) {

	u := &User{}
	u.A = 10
	fmt.Println(u)
	//user1 := User{
	//	A: 1,
	//	B: "daw",
	//	C: 7,
	//}
	//user2 := User{
	//	A: 12,
	//	B: "dadadwadaww",
	//}
	////user2 = ObjCopy(user1, user2).(User)
	//ObjCopy(&user1, &user2)
	//
	//fmt.Println(user1.A)
	//fmt.Println(user1.B)
	//fmt.Println(user1.C)
	//fmt.Println(user2.A)
	//fmt.Println(user2.B)
	//fmt.Println(user2.C)
}
