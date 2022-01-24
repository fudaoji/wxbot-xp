package core

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

var (
	WsConfig wsConfig
)

// Redis配置
type wsConfig struct {
	Host   string // 主机
	Port   string // 端口
	Path   string //
	Schema string //
}

// InitWsConfig 初始化Websocket配置
func InitWsConfig() {
	//主机
	host := GetVal("ws.host", "127.0.0.1")
	// 端口
	port := GetVal("ws.port", "5555")

	schema := GetVal("ws.schema", "")

	path := GetVal("ws.path", "/")

	WsConfig = wsConfig{
		Host:   host,
		Port:   port,
		Path:   path,
		Schema: schema,
	}
}

//InitConfig 读取配置文件
func InitConfig() {
	mode := flag.String("mode", "dev", "dev mode")
	flag.Parse()
	viper.SetConfigFile(fmt.Sprintf("./setting_%s.yaml", *mode))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

//GetVal 获取配置文件的字符串配置值
func GetVal(key string, defaultVal string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}
	return defaultVal
}

//GetIntVal 获取配置文件的整型配置值
func GetIntVal(key string, defaultVal int) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}
	return defaultVal
}
