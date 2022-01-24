package bot

import (
	"fmt"
	"net/url"
	"os"
	"wxbot-xp/core"

	"github.com/gorilla/websocket"
)

var sockConn *websocket.Conn

// public
func GetInstance() *websocket.Conn {
	if sockConn == nil {
		core.InitWsConfig()
		config := core.WsConfig
		// init
		// schema – can be ws
		// host, port – WebSocket server
		u := url.URL{
			Scheme: config.Schema,
			Host:   config.Host + ":" + config.Port,
			Path:   config.Path,
		}
		sockConn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "connect err: %v", err)
		}
		return sockConn
	}
	return sockConn
}

func InitWSConnHandle() {
	sockConn := GetInstance()
	defer sockConn.Close()
	fmt.Println("ws connect success!")

	for {
		ReceiveHandle(sockConn)
	}
}
