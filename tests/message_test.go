package tests

import (
	"fmt"
	"testing"
)

// 发送消息请求体
type sendMsg struct {
	// 送达人UserName
	To string
	// 正文
	Content string
}

//发送消息给好友
func TestTextToFriend(t *testing.T) {
	method, url := "POST", Apis["msgToFriend"]
	res := Request(method, url, sendMsg{To: "wxid_xokb2ezu1p6t21", Content: "test"})
	fmt.Printf("%#v", res)
}

func TestImgToFriend(t *testing.T) {
	method, url := "POST", Apis["imgToFriend"]
	res := Request(method, url, sendMsg{To: "wxid_xokb2ezu1p6t21", Content: "C:\\Users\\Administrator\\Documents\\WeChat Files\\wxid_a98qqf9m4bny22\\FileStorage\\Image\\2022-01\\2d4ea46c94e7ae96753f6af661da136c.dat"})
	fmt.Printf("%#v", res)
}
