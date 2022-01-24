package bot

import (
	"fmt"
	"net/url"
	"os"
	"wxbot-xp/core"

	"github.com/gorilla/websocket"
)

var SockConn *websocket.Conn

func InitWSConnHandle() {
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
	SockConn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect err: %v", err)
	}
	defer SockConn.Close()
	fmt.Println("ws connect success!")

	for {
		ReceiveHandle(SockConn)
	}
}
