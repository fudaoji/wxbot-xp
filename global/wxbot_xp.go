package global

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
	"wxbot-xp/logger"

	"github.com/gorilla/websocket"
)

const (
	HEART_BEAT              = 5005
	RECV_TXT_MSG            = 1
	RECV_PIC_MSG            = 3
	USER_LIST               = 5000
	GET_USER_LIST_SUCCSESS  = 5001
	GET_USER_LIST_FAIL      = 5002
	TXT_MSG                 = 555
	PIC_MSG                 = 500
	AT_MSG                  = 550
	CHATROOM_MEMBER         = 5010
	CHATROOM_MEMBER_NICK    = 5020
	PERSONAL_INFO           = 6500
	DEBUG_SWITCH            = 6000
	PERSONAL_DETAIL         = 6550
	DESTROY_ALL             = 9999
	NEW_FRIEND_REQUEST      = 37    //微信好友请求消息
	AGREE_TO_FRIEND_REQUEST = 10000 //同意微信好友请求消息
	ATTATCH_FILE            = 5003
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
	Mtype   uint        `json:"type"`
	Wxid    string      `json:"wxid"`
	Content interface{} `json:"content"`
	Time    string      `json:"time"`
}

const (
	Scheme string = "ws"
	Host   string = "124.223.70.93"
	Port   string = "5555"
	Path   string = "/"
)

var SockConn *websocket.Conn

func getId() string {
	return string(time.Now().Unix())
}

func getPersonalInfo() []byte {
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

func SendTxtMsg(wxid string, content string) []byte {
	j := &Message{
		Id:       getId(),
		Mtype:    TXT_MSG,
		Wxid:     wxid,
		Roomid:   "",
		Content:  content,
		Nickname: "",
		Ext:      "",
	}
	msg, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return msg
}

func receiveHandler(conn *websocket.Conn) {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error in receive:", err)
		return
	}
	log.Printf("Received: %s\n", msg)
	var data RecvMessage
	if err := json.Unmarshal(msg, &data); err != nil {
		logger.Log.Errorf("Unmarshal msg error:%v", err.Error())
	}

	switch data.Mtype {
	case RECV_TXT_MSG:
		if data.Content == "微笑" {
			err = conn.WriteMessage(websocket.TextMessage, SendTxtMsg(data.Wxid, "[微笑]"))
		}
	case RECV_PIC_MSG:

	}
	if err != nil {
		logger.Log.Errorf("send msg error:%v", err.Error())
	}
}

func InitWSConnHandle() {
	// init
	// schema – can be ws
	// host, port – WebSocket server
	u := url.URL{
		Scheme: Scheme,
		//Host:   "127.0.0.1:5555",
		Host: Host + ":" + Port,
		Path: Path,
	}
	SockConn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect err: %v", err)
	}
	defer SockConn.Close()
	fmt.Println("connect success")

	for {
		receiveHandler(SockConn)
	}
}
