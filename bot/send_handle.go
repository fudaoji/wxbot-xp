package bot

import (
	"encoding/json"
	"strings"
	"time"
)

type Message struct {
	Id       string `json:"id"`
	Mtype    uint   `json:"type"`
	Wxid     string `json:"wxid"`
	Roomid   string `json:"roomid"`
	Content  string `json:"content"`
	Nickname string `json:"nickname"`
	Ext      string `json:"ext"`
}

func getId() string {
	return string(time.Now().Unix())
}

func GetPersonalInfo() []byte {
	j := &Message{
		Id:       getId(),
		Mtype:    TXT_MSG,
		Wxid:     "",
		Roomid:   "",
		Content:  " a new socket has connected.",
		Nickname: "doogie",
		Ext:      "",
	}
	msg, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return msg
}

//发送文件消息
func SendFileMsg(wxid string, content string) []byte {
	j := &Message{
		Id:       getId(),
		Mtype:    ATTATCH_FILE,
		Wxid:     "",
		Roomid:   "",
		Content:  content,
		Nickname: "",
		Ext:      "",
	}
	if strings.Contains(wxid, "@") {
		j.Roomid = wxid
	} else {
		j.Wxid = wxid
	}
	msg, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return msg
}

//发送图片消息
func SendImg(wxid string, content string) []byte {
	j := &Message{
		Id:       getId(),
		Mtype:    PIC_MSG,
		Wxid:     "",
		Roomid:   "",
		Content:  content,
		Nickname: "",
		Ext:      "",
	}
	if strings.Contains(wxid, "@") {
		j.Roomid = wxid
	} else {
		j.Wxid = wxid
	}
	msg, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return msg
}

//发送文本消息
func SendTxt(wxid string, content string) []byte {
	j := &Message{
		Id:       getId(),
		Mtype:    TXT_MSG,
		Wxid:     "",
		Roomid:   "",
		Content:  content,
		Nickname: "",
		Ext:      "",
	}
	if strings.Contains(wxid, "@") {
		j.Roomid = wxid
	} else {
		j.Wxid = wxid
	}
	msg, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return msg
}