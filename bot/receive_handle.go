package bot

import (
	"encoding/json"
	"log"
	"reflect"
	"wxbot-xp/core"
	. "wxbot-xp/core"
	"wxbot-xp/logger"

	"github.com/gorilla/websocket"
)

/* {
	"content":"test",
	"id":"20220123165737",
	"id1":"",
	"id2":"wxid_xokb2ezu1p6t21",
	"id3":"<msgsource>\n\t<signature>v1_9mQn6Wvc</signature>\n</msgsource>\n",
	"srvid":1,
	"time":"2022-01-23 16:57:37",
	"type":1,
	"wxid":"wxid_xokb2ezu1p6t21"
} */
type RecvMessage struct {
	Id      string      `json:"id"`
	Id1     string      `json:"id1"` //发信人ID，同wxid
	Id2     string      `json:"id2"` //群ID
	Id3     string      `json:"id3"`
	Srvid   uint        `json:"srvid"`
	Mtype   int         `json:"type"`
	Wxid    string      `json:"wxid"`
	Content interface{} `json:"content"`
	Time    string      `json:"time"`
}

//好友信息结构体
type RespUserInfo struct {
	Headimgurl string `json:"headimgurl"`
	Nickname   string `json:"nickname"`
	Remarkname string `json:"remarkname"`
	Wxcode     string `json:"wxcode"`
	Wxid       string `json:"wxid"`
}

// 回调请求体
type CallbackRes struct {
	Appkey string `json:"appkey"`
	RecvMessage
}

//ReceiveHandle 消息接收器
func ReceiveHandle(conn *websocket.Conn) {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error in receive:", err.Error())
		return
	}

	var data RecvMessage
	if err := json.Unmarshal(msg, &data); err != nil {
		logger.Log.Errorf("Unmarshal msg error:%v", err.Error())
		return
	}

	log.Printf("Receive: %v", data)
	switch data.Mtype {
	case RECV_TXT_MSG:
		fallthrough
	case RECV_PIC_MSG:
		resp := CallbackRes{RecvMessage: data, Appkey: ""}
		NotifyWebhook(&resp)
	case USER_LIST:
		/* var userList []RespUserInfo
		for _, user := range data.Content {
			userList = append(userList, RespUserInfo{
				Headimgurl: user.headimg,
				Nickname:   user.name,
				Remarkname: user.remarks,
				Wxcode:     user.wxcode,
				Wxid:       user.wxid,
			})
		} */
		log.Printf("Success: %v", InterfToSlice(data.Content)...)
	case GET_USER_LIST_SUCCSESS:
		log.Printf("Success: %v", data)
	case GET_USER_LIST_FAIL:
		log.Printf("Fail: %v", data)
	}
}

//NotifyWebhook  通知客户端平台
func NotifyWebhook(data *CallbackRes) {
	url := core.GetVal("webhook", "")
	if len(url) > 1 {
		data.Appkey = core.GetVal("appkey", "")
		ReqPostJson(url, data, nil)
	}
}

//InterfToSlice 接口转为slice
func InterfToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("arr is not a slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
